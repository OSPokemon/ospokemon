package main

import (
	"github.com/ospokemon/ospokemon"
	"github.com/ospokemon/ospokemon/persistence"
	"github.com/ospokemon/ospokemon/run"
	"github.com/ospokemon/ospokemon/server"
	"taylz.io/db/patch"
	"taylz.io/log"
	"taylz.io/types"
)

const PATCH = 9

func main() {
	env := ospokemon.ENV().ParseDefault()
	logf := log.DefaultFormatWithColor()
	loglvl, err := log.GetLevel(env["loglevel"])
	if err != nil {
		log.StdOutService(log.LevelDebug, logf).Error(`env["loglevel"]`, env["loglevel"], err)
		return
	}
	logp := env["logpath"]
	ospokemon.SetLogger(log.NewService(loglvl, logf, log.NewRoller(logp)))
	log := ospokemon.LOG()
	log.Add("Patch", PATCH).Debug("ospokemon: starting...")

	persistence.OpenEnv(env)
	go run.Run(env)

	if patch, err := patch.Get(persistence.Connection); err != nil {
		log.Error("Failed to open database")
		return
	} else if patch != PATCH {
		log.Add("Found", patch).Add("Expected", PATCH).Error("Database patch mismatch")
		return
	}

	if env["edit"] != "" {
		log.Info("starting edit...")
		Editor()
		log.Info("edit finished...")
		return
	}

	log.With(types.Dict{
		"loglevel": loglvl,
		"port":     env["port"],
	}).Info("OSPokemon Server")
	server.LaunchEnv(env)
}
