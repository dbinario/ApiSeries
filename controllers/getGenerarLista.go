package controllers

import (
	"apiseries/app"
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func GetGenerarLista(c *gin.Context) {

	var db *sql.DB

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

	fmt.Println(rangos)

	//llenamos la lista de acuerdo al numeor de episodios pendientes por ver por temporadas activas

	for campo, valor := range rangos {

		for i := 0; i < valor; i++ {

			lista = append(lista, campo)

		}

	}

	fmt.Println(lista)

	// Establecer la semilla del generador de números aleatorios
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Mezclar los números de manera aleatoria
	for i := range lista {
		j := r.Intn(i + 1)
		lista[i], lista[j] = lista[j], lista[i]
	}

	//en la lista ya el orden de ver las series

	fmt.Println(lista)

	//extraemos la informacion

	// Realizar la consulta
	rows, err = db.Query("SELECT id_serie,numero_temporada,nombre_temporada FROM temporadas WHERE estado_temporada=1")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
}
