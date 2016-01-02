package loader

import (
	"github.com/ospokemon/ospokemon/objects/entities"
	"github.com/ospokemon/ospokemon/registry"
	"github.com/ospokemon/ospokemon/server"
	"github.com/ospokemon/ospokemon/world"
	"log"
)

func ConnectClient(client *server.Client) {
	session := server.Sessions[client.SessionId]
	account := server.Accounts[session.Username]

	player := LoadPlayer(account.PlayerId)
	world.AddEntity(player)
	client.AddEntity(player.EntityId())

	for _, pokemonId := range player.Pokemon() {
		LoadPokemon(pokemonId)
	}

	player.On("SummonPokemon", client.AddEntity)
	player.On("DismissPokemon", client.RemoveEntity)
}

func DisconnectClient(client *server.Client) {
	log.Println("DisconnectClient")
	session := server.Sessions[client.SessionId]
	account := server.Accounts[session.Username]
	player := registry.Players[account.PlayerId]

	client.RemoveEntity(player.EntityId())
	world.RemoveEntity(player.EntityId())

	for _, entityId := range client.Entities {
		pokemonId := world.Entities[entityId].(*entities.PokemonEntity).Id()

		if registry.Pokemon[pokemonId].EntityId() > 0 {
			world.RemoveEntity(entityId)
			UnloadPokemon(registry.Pokemon[pokemonId])
		}
	}

	UnloadPlayer(player)
}
