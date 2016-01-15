package linker

import (
	"github.com/ospokemon/ospokemon/engine"
	"github.com/ospokemon/ospokemon/objects"
	"github.com/ospokemon/ospokemon/server"
)

func ConnectClient(client *server.Client) {
	sessionId := client.SessionId
	session := server.Sessions[sessionId]
	account := server.Accounts[session.Username]
	trainerId := account.TrainerId
	trainer := objects.GetTrainer(trainerId)
	m := engine.GetMap(*trainer.Map())
	entityId := m.AddEntity(trainer)
	client.AddEntity(entityId)
	m.AddClient(client.ClientId)
}
