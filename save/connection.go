package save

import (
	"database/sql"
	"github.com/Sirupsen/logrus"
	_ "github.com/mattn/go-sqlite3"
	"github.com/ospokemon/ospokemon/util"
)

var Connection *sql.DB

func init() {
	var err error
	Connection, err = sql.Open("sqlite3", util.Opt("dbpath"))

	if err != nil {
		logrus.Error(err)
	}
}
