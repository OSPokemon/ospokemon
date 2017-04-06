package main

import (
	"ospokemon.com/log"
	_ "ospokemon.com/option"
	"ospokemon.com/persistence"
	"ospokemon.com/run"
	"ospokemon.com/server"
)

const PATCH uint64 = 9

func main() {
	log.Add("Patch", PATCH).Info("OSPokemon")

	persistence.Patch()

	if patch := persistence.CheckPatch(); patch != PATCH {
		log.Add("Found", patch).Add("Expected", PATCH).Error("Database patch mismatch")
		return
	}

	go run.Run()
	server.Launch()
}
