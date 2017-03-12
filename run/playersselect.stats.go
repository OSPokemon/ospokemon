package run

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.PlayersSelect, PlayersSelectStats)
}

func PlayersSelectStats(args ...interface{}) {
	player := args[0].(*game.Player)

	stats, err := query.PlayersStatsSelect(player)

	if err != nil {
		log.Add("Username", "2").Add("Error", err.Error()).Error("players select stats")
		return
	}

	player.AddPart(stats)
}
