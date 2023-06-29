package controllers

import (
	"apiseries/app"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func GetGenerarLista(c *gin.Context) {

	var db *sql.DB

	//se obtiene la conexion de la base de datos

	app.Setup()
	db = app.GetDB()
	defer db.Close()

}
