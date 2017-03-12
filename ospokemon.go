package main

import (
	"github.com/ospokemon/ospokemon/log"
	_ "github.com/ospokemon/ospokemon/option"
	"github.com/ospokemon/ospokemon/query"
	"github.com/ospokemon/ospokemon/run"
	"github.com/ospokemon/ospokemon/server"
)

const PATCH uint64 = 8

func main() {
	log.Add("Patch", PATCH).Info("OSPokemon")

	query.Patch()

	if patch := query.CheckPatch(); patch != PATCH {
		log.Add("Found", patch).Add("Expected", PATCH).Error("Database patch mismatch")
		return
	}

	go run.Run()
	server.Launch()
}
