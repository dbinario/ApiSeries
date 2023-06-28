package controllers

import (
	"apiseries/app"
	"apiseries/interfaces"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostEpisodioVisto(c *gin.Context) {

	var db *sql.DB

	//se obtiene la conexion de la base de datos

	app.Setup()
	db = app.GetDB()
	defer db.Close()

	var episodioVisto interfaces.EpisodioVisto

	c.ShouldBindJSON(&episodioVisto)

	// actualizamos el estado del episodio
	stmt, err := db.Prepare("UPDATE episodios SET estado='1',fecha_estado=NOW() WHERE id_serie=? AND numero_temporada=? AND numero_episodio=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	// Ejecutar la consulta con los parámetros
	_, err = stmt.Exec(episodioVisto.IdSerie, episodioVisto.Temporada, episodioVisto.Episodio)
	if err != nil {
		panic(err.Error())
	}

	// actualizamos el estado del episodio
	stmt, err = db.Prepare("UPDATE temporadas SET ultimo_visto=? WHERE id_serie=? AND numero_temporada=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	// Ejecutar la consulta con los parámetros
	_, err = stmt.Exec(episodioVisto.Episodio, episodioVisto.IdSerie, episodioVisto.Temporada)
	if err != nil {
		panic(err.Error())
	}

	var nombreSerie string
	var nombreEpisodio string
	var nombreTemporada string

	db.QueryRow("SELECT nombre_episodio FROM episodios WHERE numero_episodio=? AND id_serie=? AND numero_temporada=?", episodioVisto.Episodio, episodioVisto.IdSerie, episodioVisto.Temporada).Scan(&nombreEpisodio)
	db.QueryRow("SELECT nombre_serie FROM series WHERE id_serie=?", episodioVisto.IdSerie).Scan(&nombreSerie)
	db.QueryRow("SELECT nombre_temporada FROM temporadas WHERE numero_temporada=? AND id_serie=?", episodioVisto.Temporada, episodioVisto.IdSerie).Scan(&nombreTemporada)

	db.Exec("CALL actualizarVistos();")

	c.JSON(http.StatusOK, gin.H{
		"message": "El capitulo: " + nombreEpisodio + " de la temporada: " + nombreTemporada + " de la serie " + nombreSerie + " se ha actualizado",
	})

}
