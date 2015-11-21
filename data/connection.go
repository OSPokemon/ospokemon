package data

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var Connection *sql.DB

func Connect(database string) {
	var err error
	Connection, err = sql.Open("sqlite3", database)

	if err != nil {
		log.Fatal(err)
	}
}
