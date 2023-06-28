package controllers

import (
	"apiseries/app"
	"apiseries/interfaces"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteSerie(c *gin.Context) {

	var db *sql.DB

	//se obtiene la conexion de la base de datos

	app.Setup()
	db = app.GetDB()
	defer db.Close()

	var eliminarSerie interfaces.EliminarSerie

	c.ShouldBindJSON(&eliminarSerie)

	var nombreserie string

	db.QueryRow("SELECT nombre_serie FROM series WHERE id_serie=?", eliminarSerie.IdSerie).Scan(&nombreserie)

	// eliminamos la serie completa
	stmt, err := db.Prepare("CALL eliminarSerie(?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	// Ejecutar la consulta con los parámetros
	_, err = stmt.Exec(eliminarSerie.IdSerie)
	if err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "La serie: " + nombreserie + " se elimino de manera correcta",
	})

}
