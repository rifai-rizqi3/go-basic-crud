package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "rnrifai:RNRif@i1212@/go_products")
	if err != nil {
		panic(err)
	}

	log.Println("Connected to database")
	DB = db
}
