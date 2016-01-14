package main

import (
	"github.com/ospokemon/ospokemon/db"
	"github.com/ospokemon/ospokemon/engine"
	"github.com/ospokemon/ospokemon/linker"
	"github.com/ospokemon/ospokemon/objects"
	"github.com/ospokemon/ospokemon/server"
)

func wiring() {
	server.CreateAccount = db.CreateAccount
	server.LoadAccount = db.LoadAccount
	server.ChangePassword = db.ChangePassword
	server.ConnectClient = linker.ConnectClient
	server.ReceiveMessage = engine.ReceiveMessage
	objects.GetSpeciesIds = db.GetSpeciesIds
	objects.LoadSpecies = db.LoadSpecies
	objects.LoadTrainer = db.LoadTrainer
	objects.CreateTrainer = db.CreateTrainer
	objects.MakePokemon = linker.MakePokemon
	objects.SavePokemon = db.SavePokemon
	objects.LoadClass = db.LoadClass
	objects.GetClassIds = db.GetClassIds
}
