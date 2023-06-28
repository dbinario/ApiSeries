package controllers

import (
	"apiseries/app"
	"apiseries/interfaces"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostCapituloVisto(c *gin.Context) {

	var db *sql.DB

	//se obtiene la conexion de la base de datos

	app.Setup()
	db = app.GetDB()
	defer db.Close()

	var capitulovisto interfaces.CapituloVisto

	c.ShouldBindJSON(&capitulovisto)

	// actualizamos el estado del capitulo
	stmt, err := db.Prepare("UPDATE capitulos SET estado='1',fecha_estado=NOW() WHERE id_serie=? AND numero_temporada=? AND numero_episodio=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	// Ejecutar la consulta con los parámetros
	_, err = stmt.Exec(capitulovisto.IdSerie, capitulovisto.Temporada, capitulovisto.Capitulo)
	if err != nil {
		panic(err.Error())
	}

	var nombreserie string
	var nombrecapitulo string
	var nombretemporada string

	db.QueryRow("SELECT nombre_capitulo FROM capitulos WHERE numero_episodio=? AND id_serie=? AND numero_temporada=?", capitulovisto.Capitulo, capitulovisto.IdSerie, capitulovisto.Temporada).Scan(&nombrecapitulo)
	db.QueryRow("SELECT nombre_serie FROM series WHERE id_serie=?", capitulovisto.IdSerie).Scan(&nombreserie)
	db.QueryRow("SELECT nombre_temporada FROM temporadas WHERE numero_temporada=? AND id_serie=?", capitulovisto.Temporada, capitulovisto.IdSerie).Scan(&nombretemporada)

	db.Exec("CALL actualizarVistos();")

	c.JSON(http.StatusOK, gin.H{
		"message": "El capitulo: " + nombrecapitulo + " de la temporada: " + nombretemporada + " de la serie " + nombreserie + " se ha actualizado",
	})

}
