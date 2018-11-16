package main // import "ospokemon.com/cmd/ospokemon"

import (
	"ospokemon.com/persistence"
	"ospokemon.com/run"
	"ospokemon.com/server"
	"ztaylor.me/db"
	"ztaylor.me/env"
	"ztaylor.me/log"
)

const PATCH = 9

func main() {
	env := env.Global()
	log.SetLevel(env.Default("loglevel", "info"))
	log.Add("Patch", PATCH).Debug("ospokemon: starting...")

	persistence.OpenEnv(env)
	go run.Run(env)

	if patch, err := db.Patch(persistence.Connection); err != nil {
		log.Error("Failed to open database")
		return
	} else if patch != PATCH {
		log.Add("Found", patch).Add("Expected", PATCH).Error("Database patch mismatch")
		return
	}

	if env.Get("edit") != "" {
		log.Info("starting edit...")
		Editor()
		log.Info("edit finished...")
		return
	}

	log.WithFields(log.Fields{
		"loglevel": env.Get("loglevel"),
		"port":     env.Get("port"),
	}).Info("OSPokemon Server")
	server.LaunchEnv(env)
}
