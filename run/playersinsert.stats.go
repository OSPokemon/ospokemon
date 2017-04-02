package run

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
	"ospokemon.com/persistence"
)

func init() {
	event.On(event.PlayersInsert, PlayersInsertStats)
}

func PlayersInsertStats(args ...interface{}) {
	player := args[0].(*ospokemon.Player)
	stats := player.GetStats()

	if len(stats) < 1 {
		stats = map[string]*ospokemon.Stat{
			"speed": &ospokemon.Stat{5, 5, 5},
		}

		log.Add("Username", player.Username).Debug("players insert stats: grant default stats")
	}

	err := persistence.PlayersStatsInsert(player, stats)

	if err != nil {
		log.Add("Username", player.Username).Add("Stats", stats).Add("Error", err.Error()).Error("players insert stats")
	}
}
