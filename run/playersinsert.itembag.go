package run

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.PlayersInsert, PlayersInsertItembag)
}

func PlayersInsertItembag(args ...interface{}) {
	player := args[0].(*game.Player)
	itembag := player.GetItembag()

	if itembag == nil {
		itembag = game.MakeItembag(player.BagSize)
		logrus.WithFields(logrus.Fields{
			"Username": player.Username,
		}).Debug("players insert itembag: grant empty bag")
	}

	err := query.ItembagsPlayersInsert(player, itembag)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Player":  player.Username,
			"Itembag": itembag.GetItems(),
			"Error":   err.Error(),
		}).Error("playersinsert.itembag")
	}
}
