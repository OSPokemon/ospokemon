package persistence

import (
	"database/sql"

	"ztaylor.me/db"
	"ztaylor.me/env"
	"ztaylor.me/log"
)

var Connection *sql.DB

var Open = db.Open

func OpenEnv(env env.Provider) {
	var err error
	Connection, err = db.OpenEnv(env)

	log := log.Add("Host", env.Get(db.DB_HOST))

	if err != nil {
		log.Add("Error", err.Error()).Error("persistence.Connect")
	} else {
		log.Info("persistence.Connect")
	}
}
