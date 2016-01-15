package linker

import (
	"github.com/ospokemon/ospokemon/engine"
	// "github.com/ospokemon/ospokemon/objects"
	"github.com/ospokemon/ospokemon/server"
)

func DisconnectClient(client *server.Client) {
	for _, entityId := range client.Entities {
		entity := engine.Entities[entityId]
		mapId := *entity.Map()
		m := engine.Maps[mapId]
		m.RemoveEntity(entityId)
		m.RemoveClient(client.ClientId)
	}
}
