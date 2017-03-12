package query

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/ospokemon/ospokemon/log"
	"github.com/ospokemon/ospokemon/option"
)

var Connection *sql.DB

func init() {
	var err error
	Connection, err = sql.Open("sqlite3", option.String("dbpath"))

	if err != nil {
		log.Add("Path", option.String("dbpath")).Add("Error", err.Error()).Error("query.init")
	}
}
