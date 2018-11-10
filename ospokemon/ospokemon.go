package main

import (
	"ospokemon.com/persistence"
	"ospokemon.com/run"
	"ospokemon.com/server"
	"ztaylor.me/db"
	"ztaylor.me/env"
	"ztaylor.me/log"
)

const PATCH = 12

func main() {
	env := env.Global()
	log.Add("Patch", PATCH).Info("OSPokemon")

	persistence.Connect(env)

	if patch, err := db.Patch(persistence.Connection); err != nil {
		log.Error("Failed to open database")
		return
	} else if patch != PATCH {
		log.Add("Found", patch).Add("Expected", PATCH).Error("Database patch mismatch")
		return
	}

	go run.Run(env)
	server.Launch(env)
}
