package main

import (
	"github.com/ospokemon/ospokemon/server"
	"github.com/ospokemon/ospokemon/util"
)

const APP_VERSION = "0.0"

func main() {
	util.Log.Info("OSPokemon ", APP_VERSION)
	util.LogOptions()
	server.Launch()
}