package save

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/ospokemon/ospokemon/util"
)

var Connection *sql.DB

func Connect(database string) {
	var err error
	Connection, err = sql.Open("sqlite3", util.Opt("dbpath"))

	if err != nil {
		util.Log.Error(err)
	}
}
