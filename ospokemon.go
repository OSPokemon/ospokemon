package main

import (
	"github.com/Sirupsen/logrus"
	_ "github.com/ospokemon/ospokemon/cmd"
	"github.com/ospokemon/ospokemon/save"
	"github.com/ospokemon/ospokemon/server"
	"github.com/ospokemon/ospokemon/util"
)

const PATCH uint64 = 4

func main() {
	logrus.WithFields(logrus.Fields{
		"Patch": PATCH,
	}).Info("OSPokemon")

	if util.Opt("patchpath") != "" {
		save.Patch()
		return
	}

	if patch := save.CheckPatch(); patch != PATCH {
		logrus.WithFields(logrus.Fields{
			"Found":    patch,
			"Expected": PATCH,
		}).Fatal("save.CheckPatch: Database patch mismatch")
		return
	}

	server.Launch()
}
