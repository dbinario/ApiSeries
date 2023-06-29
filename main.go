package main

import (
	"apiseries/controllers"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando archivo .env")
	}

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/series", controllers.GetSeries)
	router.GET("/queseriever", controllers.GetQueSerieVer)
	router.GET("/generarlista", controllers.GetGenerarLista)
	router.POST("/serie", controllers.PostRecuperarSerie)
	router.POST("/capitulovisto", controllers.PostEpisodioVisto)
	router.DELETE("/eliminarserie", controllers.DeleteSerie)
	router.DELETE("/eliminartemporada", controllers.DeleteTemporada)

	router.Run("localhost:8080")

}
