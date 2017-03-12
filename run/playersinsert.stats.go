package run

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.PlayersInsert, PlayersInsertStats)
}

func PlayersInsertStats(args ...interface{}) {
	player := args[0].(*game.Player)
	stats := player.GetStats()

	if len(stats) < 1 {
		stats = map[string]*game.Stat{
			"speed": &game.Stat{5, 5, 5},
		}

		logrus.WithFields(logrus.Fields{
			"Username": player.Username,
		}).Debug("players insert stats: grant default stats")
	}

	err := query.PlayersStatsInsert(player, stats)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Username": player.Username,
			"Stats":    stats,
			"Error":    err.Error(),
		}).Error("players insert stats")
	}
}
