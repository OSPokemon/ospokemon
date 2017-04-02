package persistence

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"ospokemon.com/log"
	"ospokemon.com/option"
)

var Connection *sql.DB

func init() {
	var err error
	Connection, err = sql.Open("sqlite3", option.String("dbpath"))

	if err != nil {
		log.Add("Path", option.String("dbpath")).Add("Error", err.Error()).Error("query.init")
	}
}
