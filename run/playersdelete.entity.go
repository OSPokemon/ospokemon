package run

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/part"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.PlayersDelete, PlayersDeleteEntity)
}

func PlayersDeleteEntity(args ...interface{}) {
	player := args[0].(*game.Player)
	entity := player.Parts[part.Entity].(*game.Entity)

	universe := game.Multiverse[entity.UniverseId]
	universe.Remove(entity)

	err := query.EntitiesPlayersDelete(player)

	if err != nil {
		logrus.Error(err.Error())
	}
}
