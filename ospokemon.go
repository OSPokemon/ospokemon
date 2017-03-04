package main

import (
	"github.com/Sirupsen/logrus"
	_ "github.com/ospokemon/ospokemon/option"
	"github.com/ospokemon/ospokemon/query"
	"github.com/ospokemon/ospokemon/run"
	"github.com/ospokemon/ospokemon/server"
)

const PATCH uint64 = 8

func main() {
	logrus.WithFields(logrus.Fields{
		"Patch": PATCH,
	}).Info("OSPokemon")

	query.Patch()

	if patch := query.CheckPatch(); patch != PATCH {
		logrus.WithFields(logrus.Fields{
			"Found":    patch,
			"Expected": PATCH,
		}).Fatal("Database patch mismatch")
		return
	}

	go run.Run()
	server.Launch()
}
