package run

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.PlayersDelete, PlayersDeleteStats)
}

func PlayersDeleteStats(args ...interface{}) {
	player := args[0].(*game.Player)
	err := query.PlayersStatsDelete(player)

	if err != nil {
		logrus.Error(err.Error())
	}
}
