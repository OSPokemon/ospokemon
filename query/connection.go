package query

import (
	"database/sql"
	"github.com/Sirupsen/logrus"
	_ "github.com/mattn/go-sqlite3"
	"github.com/ospokemon/ospokemon/option"
)

var Connection *sql.DB

func init() {
	var err error
	Connection, err = sql.Open("sqlite3", option.String("dbpath"))

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Path": option.String("dbpath"),
		}).Error(err)
	}
}
