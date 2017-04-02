package run

import (
	"ospokemon.com"
	"ospokemon.com/event"
)

func init() {
	event.On(event.PlayersDelete, PlayersDeleteEntity)
}

func PlayersDeleteEntity(args ...interface{}) {
	player := args[0].(*ospokemon.Player)
	entity := player.GetEntity()
	universe, _ := ospokemon.GetUniverse(entity.UniverseId)
	universe.Remove(entity)
}
