package run

import (
	"ospokemon.com/event"
	"ospokemon.com/game"
	"ospokemon.com/query"
)

func init() {
	event.On(event.PlayersDelete, PlayersDeleteEntity)
}

func PlayersDeleteEntity(args ...interface{}) {
	player := args[0].(*game.Player)
	entity := player.GetEntity()
	universe, _ := query.GetUniverse(entity.UniverseId)
	universe.Remove(entity)
}
