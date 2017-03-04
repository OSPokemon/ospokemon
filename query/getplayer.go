package query

import (
	"github.com/ospokemon/ospokemon/game"
)

func GetPlayer(username string) (*game.Player, error) {
	if player, ok := game.Players[username]; ok {
		return player, nil
	}

	return PlayersSelect(username)
}
