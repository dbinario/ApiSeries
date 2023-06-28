package app

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Setup() {

	d, err := sql.Open(os.Getenv("DB_TYPE"), os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_DATABASE"))

	if err != nil {
		panic(err)
	}

	db = d

}

func GetDB() *sql.DB {

	return db

}
