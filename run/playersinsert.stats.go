package run

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
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

		log.Add("Username", "2").Debug("players insert stats: grant default stats")
	}

	err := query.PlayersStatsInsert(player, stats)

	if err != nil {
		log.Add("Username", player.Username).Add("Stats", stats).Add("Error", err.Error()).Error("players insert stats")
	}
}
