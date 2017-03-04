package run

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/part"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.PlayersInsert, PlayersInsertItembag)
}

func PlayersInsertItembag(args ...interface{}) {
	player := args[0].(*game.Player)
	bag, ok := player.Parts[part.Itembag].(*game.Itembag)

	if !ok {
		bag = game.MakeItembag(player.BagSize)
		logrus.WithFields(logrus.Fields{
			"Username": player.Username,
		}).Debug("players insert itembag: grant empty bag")
	}

	err := query.ItembagsPlayersInsert(player, bag)

	if err != nil {
		logrus.Error(err.Error())
	}
}
