package persistence

import (
	"github.com/ospokemon/ospokemon"
	"taylz.io/db"
	"taylz.io/db/mysql"
	"taylz.io/env"
)

var Connection *db.DB

func OpenEnv(env env.Service) {
	var err error
	Connection, err = mysql.Open(db.ParseDSN(env))

	log := ospokemon.LOG().Add("Host", ospokemon.ENV()["DB_HOST"])

	if err != nil {
		log.Add("Error", err.Error()).Error("persistence.Connect")
	} else {
		log.Info("persistence.Connect")
	}
}
