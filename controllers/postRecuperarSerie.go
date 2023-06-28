package controllers

import (
	"apiseries/app"
	"apiseries/interfaces"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PostRecuperarSerie(c *gin.Context) {

	var db *sql.DB

	//se obtiene la conexion de la base de datos

	app.Setup()
	db = app.GetDB()
	defer db.Close()

	var recuperar interfaces.Recuperar
	var serie interfaces.Serie

	//mapeamos la respuesta
	c.ShouldBindJSON(&recuperar)

	//primero verificamos si ya tenemos la serie en la base de datos

	var buscarserie int

	db.QueryRow("SELECT COUNT(*) FROM series WHERE id_serie=?", recuperar.IdSerie).Scan(&buscarserie)

	//si la serie no existe en la base de datos se recuperan los datos de la serie
	if buscarserie == 0 {

		url := "https://api.themoviedb.org/3/tv/" + strconv.Itoa(recuperar.IdSerie) + "?language=en-US"
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("accept", "application/json")
		req.Header.Add("Authorization", "Bearer "+os.Getenv("TOKEN_API"))

		res, _ := http.DefaultClient.Do(req)
		body, _ := io.ReadAll(res.Body)

		err := json.Unmarshal(body, &serie)
		if err != nil {
			fmt.Println("Error al mapear la respuesta JSON:", err)
			return
		}

		defer res.Body.Close()

		// Insertar datos en la tabla
		stmt, err := db.Prepare("INSERT INTO series(id_serie,nombre_serie,numero_temporadas,numero_episodios) VALUES(?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		defer stmt.Close()

		// Ejecutar la consulta con los parámetros
		_, err = stmt.Exec(recuperar.IdSerie, serie.Name, serie.NumberOfSeasons, serie.NumberOfEpisodes)
		if err != nil {
			panic(err.Error())
		}

		for _, valor := range serie.Seasons {

			var temporada interfaces.Temporada

			urlS := "https://api.themoviedb.org/3/tv/" + strconv.Itoa(recuperar.IdSerie) + "/season/" + strconv.Itoa(valor.SeasonNumber) + "?language=en-US"

			req, _ := http.NewRequest("GET", urlS, nil)

			req.Header.Add("accept", "application/json")
			req.Header.Add("Authorization", "Bearer "+os.Getenv("TOKEN_API"))

			res, _ := http.DefaultClient.Do(req)

			body, _ := io.ReadAll(res.Body)

			err := json.Unmarshal(body, &temporada)
			if err != nil {
				fmt.Println("Error al mapear la respuesta JSON:", err)
				return
			}
			defer res.Body.Close()

			// Insertar datos en la tabla
			stmt, err := db.Prepare("INSERT INTO temporadas(id_serie,numero_temporada,nombre_temporada,numero_episodios) VALUES(?,?,?,?)")
			if err != nil {
				panic(err.Error())
			}
			defer stmt.Close()

			// Ejecutar la consulta con los parámetros
			_, err = stmt.Exec(recuperar.IdSerie, temporada.SeasonNumber, temporada.Name, valor.EpisodeCount)
			if err != nil {
				panic(err.Error())
			}

			for _, valor := range temporada.Episodes {

				// Insertar datos en la tabla
				stmt, err := db.Prepare("INSERT INTO episodios(id_serie,numero_temporada,numero_episodio,nombre_episodio) VALUES(?,?,?,?)")
				if err != nil {
					panic(err.Error())
				}
				defer stmt.Close()

				// Ejecutar la consulta con los parámetros
				_, err = stmt.Exec(recuperar.IdSerie, temporada.SeasonNumber, valor.EpisodeNumber, valor.Name)
				if err != nil {
					panic(err.Error())
				}

			}

		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Serie " + serie.Name + " recuperada de manera exitosa",
		})

	} //TODO crear else para actualizar

	db.Exec("CALL actualizarVistos();")

}
