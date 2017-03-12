package run

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.PlayersDelete, PlayersDeleteEntity)
}

func PlayersDeleteEntity(args ...interface{}) {
	player := args[0].(*game.Player)
	entity := player.GetEntity()
	universe, err := query.GetUniverse(entity.UniverseId)
	universe.Remove(entity)

	if err != nil {
		logrus.Error(err.Error())
	}
}
