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

func GetQueSerieVer(c *gin.Context) {

	var db *sql.DB

	//se obtiene la conexion de la base de datos

	app.Setup()
	db = app.GetDB()
	defer db.Close()

	var QueSerieVer interfaces.QueSerieVer

	//aqui tengo que desarrollar el algoritmo para poder elegir al azar que serie y capitulo debo de ver

	//creamos el mapa para guardar los indices de la serie y los capitulos por ver
	rangos := make(map[int][2]int)

	// Realizar la consulta
	rows, err := db.Query("SELECT id_serie,SUM(faltan_ver) AS faltan_Ver FROM temporadas WHERE estado_temporada=1 GROUP BY id_serie")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Sumar los valores del mapa para saber cuantos capitulos tenemos pendientes por ver
	suma := 0
	rangoini := 1
	rangofin := 0
	for rows.Next() {

		var idserie int
		var capitulos int

		err := rows.Scan(&idserie, &capitulos)
		if err != nil {
			panic(err)
		}

		suma += capitulos
		rangofin += capitulos
		rangos[idserie] = [2]int{rangoini, rangofin}
		rangoini += capitulos
	}

	// Establecer la semilla del generador de números aleatorios
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generar un número aleatorio entre 1 y el valor maximo de suma
	numero := r.Intn(suma) + 1

	// Verificar en qué rango se encuentra el número aleatorio

	var serie int
	for nombre, rango := range rangos {
		if numero >= rango[0] && numero <= rango[1] {
			fmt.Println(nombre, rango[0], rango[1])
			serie = nombre
			break
		}
	}

	//aqui recuperamos los datos a mostrar

	var nombreSerie string
	var nombreTemporada string
	var numeroTemporada int
	var ultimoVisto int
	var nombreEpisodio string

	db.QueryRow("SELECT nombre_serie FROM series WHERE id_serie=?", serie).Scan(&nombreSerie)
	db.QueryRow("SELECT nombre_temporada,numero_temporada,ultimo_visto +1 as ultvis FROM temporadas WHERE id_serie=? AND estado_temporada=1", serie).Scan(&nombreTemporada, &numeroTemporada, &ultimoVisto)
	db.QueryRow("SELECT nombre_episodio FROM episodios WHERE id_serie=? AND numero_temporada=? AND numero_episodio=?", serie, numeroTemporada, ultimoVisto).Scan(&nombreEpisodio)

	//aqui llenamos los datos para el json

	QueSerieVer.IdSerie = serie
	QueSerieVer.NombreSerie = nombreSerie
	QueSerieVer.Temporada = numeroTemporada
	QueSerieVer.NombreTemporada = nombreTemporada
	QueSerieVer.Episodio = ultimoVisto
	QueSerieVer.NombreEpisodio = nombreEpisodio

	//imprimimos el resultado
	c.JSON(http.StatusOK, gin.H{
		"message": QueSerieVer,
	})

}
