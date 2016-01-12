package db

import (
	"database/sql"
	log "github.com/Sirupsen/logrus"
	_ "github.com/mattn/go-sqlite3"
)

var Connection *sql.DB

func Connect(database string) {
	var err error
	Connection, err = sql.Open("sqlite3", database)

	if err != nil {
		log.Fatal(err)
	}
}
