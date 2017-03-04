package query

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
)

func PlayersSelect(username string) (*game.Player, error) {
	row := Connection.QueryRow(
		"SELECT level, experience, money, class, bagsize FROM players WHERE username=?",
		username,
	)

	player := game.MakePlayer(username)
	err := row.Scan(&player.Level, &player.Experience, &player.Money, &player.Class, &player.BagSize)

	if err != nil {
		game.Players[username] = nil
		return nil, err
	}

	game.Players[username] = player

	logrus.WithFields(logrus.Fields{
		"Username": player.Username,
	}).Info("players select")

	event.Fire(event.PlayersSelect, player)

	return player, nil
}
