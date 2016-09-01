package main

import (
	_ "github.com/ospokemon/ospokemon/cmd"
	"github.com/ospokemon/ospokemon/save"
	"github.com/ospokemon/ospokemon/server"
	"github.com/ospokemon/ospokemon/util"
)

func main() {
	util.Log.Info("OSPokemon")

	if save.CheckPatch() {
		server.Launch()
	}
}
