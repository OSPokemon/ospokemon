package main

import (
	"github.com/ospokemon/ospokemon/save"
	"github.com/ospokemon/ospokemon/server"
	"github.com/ospokemon/ospokemon/util"
)

func main() {
	util.Log.Info("OSPokemon")
	util.LogOptions()

	if save.CheckPatch() {
		server.Launch()
	}
}
