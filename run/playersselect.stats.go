package run

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.PlayersSelect, PlayersSelectStats)
}

func PlayersSelectStats(args ...interface{}) {
	player := args[0].(*game.Player)

	stats, err := query.PlayersStatsSelect(player)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Username": player.Username,
			"Error":    err.Error(),
		}).Error("players select stats")
		return
	}

	player.AddPart(stats)
}
