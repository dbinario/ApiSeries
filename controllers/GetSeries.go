package controllers

import (
	"apiseries/app"
	"apiseries/interfaces"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSeries(c *gin.Context) {

	var db *sql.DB

	//se obtiene la conexion de la base de datos

	app.Setup()
	db = app.GetDB()
	defer db.Close()

	db.Exec("CALL actualizarVistos();")

	var estadoSeries interfaces.EstadoSeries

	db.QueryRow("SELECT COUNT(*) AS total,SUM(capitulos_vistos) AS vistos, SUM(faltan_ver) AS faltan_ver  FROM series").Scan(&estadoSeries.NumeroSeries, &estadoSeries.CapitulosVistos, &estadoSeries.CapitulosPorVer)

	// Realizar la consulta
	rows, err := db.Query("SELECT id_serie,nombre_serie, numero_capitulos,capitulos_vistos,faltan_ver FROM series")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {

		serie := struct {
			IDserie         int    `json:"id_serie"`
			NombreSerie     string `json:"nombre_serie"`
			NumeroCapitulos int    `json:"numero_capitulos"`
			CapitulosVistos int    `json:"capitulos_vistos"`
			CapitulosPorVer int    `json:"capitulos_por_ver"`
		}{}

		err := rows.Scan(&serie.IDserie, &serie.NombreSerie, &serie.NumeroCapitulos, &serie.CapitulosVistos, &serie.CapitulosPorVer)
		if err != nil {
			panic(err.Error())
		}

		estadoSeries.Series = append(estadoSeries.Series, serie)

	}

	// Verificar si ocurrieron errores durante la iteración
	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{

		"estado": estadoSeries,
	})
}
