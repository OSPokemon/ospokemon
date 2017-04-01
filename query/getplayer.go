package query

import (
	"ospokemon.com"
)

func GetPlayer(username string) (*ospokemon.Player, error) {
	if player, ok := ospokemon.Players[username]; ok {
		return player, nil
	}

	return PlayersSelect(username)
}
