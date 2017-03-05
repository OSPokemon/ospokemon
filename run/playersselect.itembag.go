package run

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.PlayersSelect, PlayersSelectItembag)
}

func PlayersSelectItembag(args ...interface{}) {
	player := args[0].(*game.Player)
	itembag, err := query.ItembagsPlayersSelect(player)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Username": player.Username,
			"Error":    err.Error(),
		}).Error("player build itembag")
		return
	}

	player.AddPart(itembag)
}
