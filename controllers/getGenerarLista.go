package controllers

import (
	"apiseries/app"
	"apiseries/interfaces"
	"database/sql"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetGenerarLista(c *gin.Context) {

	var db *sql.DB

	var listaSeries interfaces.ListaSeries

	//se obtiene la conexion de la base de datos

	app.Setup()
	db = app.GetDB()
	defer db.Close()

	// Realizar la consulta
	rows, err := db.Query("SELECT id_serie,SUM(faltan_ver) AS faltan_Ver FROM temporadas WHERE estado_temporada=1 GROUP BY id_serie")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	//creamos el mapa para guardar los indices de la serie y los capitulos por ver
	rangos := make(map[int]int)

	// Sumar los valores del mapa para saber cuantos capitulos tenemos pendientes por ver
	suma := 0
	for rows.Next() {

		var idserie int
		var capitulos int

		err := rows.Scan(&idserie, &capitulos)
		if err != nil {
			panic(err)
		}

		suma += capitulos
		rangos[idserie] = capitulos
	}

	lista := []int{}

	//llenamos la lista de acuerdo al numero de episodios pendientes por ver por temporadas activas

	for campo, valor := range rangos {

		for i := 0; i < valor; i++ {

			lista = append(lista, campo)

		}

	}

	listaSeries.EpisodiosPorVer = suma
	listaSeries.NumeroSeries = len(rangos)

	// Establecer la semilla del generador de números aleatorios
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Mezclar los números de manera aleatoria
	for i := range lista {
		j := r.Intn(i + 1)
		lista[i], lista[j] = lista[j], lista[i]
	}

	//recorremos las temporadas para saber que capitulos debo recorrer

	rows, err = db.Query("CALL listaEpisodios()")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	episodiosTotales := []struct {
		IDserie         int    `json:"id_serie"`
		NombreSerie     string `json:"nombre_serie"`
		NumeroTemporada int    `json:"numero_temporada"`
		NombreTemporada string `json:"nombre_temporada"`
		NumeroEpisodio  int    `json:"numero_episodio"`
		NombreEpisodio  string `json:"nombre_episodio"`
	}{}

	for rows.Next() {

		episodio := struct {
			IDserie         int    `json:"id_serie"`
			NombreSerie     string `json:"nombre_serie"`
			NumeroTemporada int    `json:"numero_temporada"`
			NombreTemporada string `json:"nombre_temporada"`
			NumeroEpisodio  int    `json:"numero_episodio"`
			NombreEpisodio  string `json:"nombre_episodio"`
		}{}

		err := rows.Scan(&episodio.IDserie, &episodio.NombreSerie, &episodio.NumeroTemporada, &episodio.NombreTemporada, &episodio.NumeroEpisodio, &episodio.NombreEpisodio)
		if err != nil {
			fmt.Println(err)
			return
		}

		//episodios = append(episodios, episodios)

		episodiosTotales = append(episodiosTotales, episodio)

	}

	//en episodiosTotales tenemos un array de de structs

	//metemos el ray de struct en un orden aleatorio conservando el orden de los capitulos en la liste series

	for _, valor := range lista {

		indice := 0
		bandera := false
		for j, k := range episodiosTotales {

			if valor == k.IDserie {

				listaSeries.Episodios = append(listaSeries.Episodios, k)

				bandera = true
				indice = j
				break
			}

		}

		if bandera {

			episodiosTotales = append(episodiosTotales[:indice], episodiosTotales[indice+1:]...)

		}

		bandera = false

	}

	c.JSON(http.StatusOK, gin.H{

		"message": listaSeries,
	})

}
