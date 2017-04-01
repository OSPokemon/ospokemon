package ospokemon

import (
	"ospokemon.com/log"
	_ "ospokemon.com/option"
	"ospokemon.com/query"
	"ospokemon.com/run"
	"ospokemon.com/server"
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
