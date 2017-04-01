package query

import (
	"ospokemon.com/game"
)

func GetPlayer(username string) (*game.Player, error) {
	if player, ok := game.Players[username]; ok {
		return player, nil
	}

	return PlayersSelect(username)
}
