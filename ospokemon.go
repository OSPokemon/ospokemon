package main

import (
	"github.com/Sirupsen/logrus"
	_ "github.com/ospokemon/ospokemon/option"
	"github.com/ospokemon/ospokemon/run"
	"github.com/ospokemon/ospokemon/save"
	_ "github.com/ospokemon/ospokemon/script"
	"github.com/ospokemon/ospokemon/server"
)

const PATCH uint64 = 4

func main() {
	logrus.WithFields(logrus.Fields{
		"Patch": PATCH,
	}).Info("OSPokemon")

	save.Patch()

	if patch := save.CheckPatch(); patch != PATCH {
		logrus.WithFields(logrus.Fields{
			"Found":    patch,
			"Expected": PATCH,
		}).Fatal("Database patch mismatch")
		return
	}

	go run.Run()
	server.Launch()
}
