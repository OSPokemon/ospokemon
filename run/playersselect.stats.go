package run

import (
	"ospokemon.com/event"
	"ospokemon.com/game"
	"ospokemon.com/log"
	"ospokemon.com/query"
)

func init() {
	event.On(event.PlayersSelect, PlayersSelectStats)
}

func PlayersSelectStats(args ...interface{}) {
	player := args[0].(*game.Player)

	stats, err := query.PlayersStatsSelect(player)

	if err != nil {
		log.Add("Username", player.Username).Add("Error", err.Error()).Error("players select stats")
		return
	}

	player.AddPart(stats)
}
