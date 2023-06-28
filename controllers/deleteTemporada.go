package controllers

import (
	"apiseries/app"
	"apiseries/interfaces"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteTemporada(c *gin.Context) {

	var db *sql.DB

	//se obtiene la conexion de la base de datos

	app.Setup()
	db = app.GetDB()
	defer db.Close()

	var eliminarTemporada interfaces.EliminarTemporada

	c.ShouldBindJSON(&eliminarTemporada)

	var nombreserie string
	var nombretemporada string

	db.QueryRow("SELECT nombre_serie FROM series WHERE id_serie=?", eliminarTemporada.IdSerie).Scan(&nombreserie)
	db.QueryRow("SELECT nombre_temporada FROM temporadas WHERE numero_temporada=? AND id_serie=?", eliminarTemporada.Temporada, eliminarTemporada.IdSerie).Scan(&nombretemporada)

	// eliminamos la serie completa
	stmt, err := db.Prepare("CALL eliminarTemporada(?,?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	// Ejecutar la consulta con los parámetros
	_, err = stmt.Exec(eliminarTemporada.IdSerie, eliminarTemporada.Temporada)
	if err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "La temporada " + nombretemporada + " de la serie: " + nombreserie + " se elimino de manera correcta",
	})

}
