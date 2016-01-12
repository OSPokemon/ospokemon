package main

import (
	"github.com/ospokemon/ospokemon/db"
	"github.com/ospokemon/ospokemon/engine"
	"github.com/ospokemon/ospokemon/linker"
	"github.com/ospokemon/ospokemon/server"
)

func wiring() {
	server.CreateAccount = db.CreateAccount
	server.LoadAccount = db.LoadAccount
	server.ChangePassword = db.ChangePassword
	server.ConnectClient = linker.ConnectClient
	server.ReceiveMessage = engine.ReceiveMessage
}
