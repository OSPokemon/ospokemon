package main

import (
	"github.com/Sirupsen/logrus"
	_ "github.com/ospokemon/ospokemon/cmd"
	"github.com/ospokemon/ospokemon/save"
	"github.com/ospokemon/ospokemon/server"
)

func main() {
	logrus.Info("OSPokemon")

	if save.CheckPatch() {
		server.Launch()
	}
}
