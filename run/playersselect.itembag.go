package run

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/part"
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

	for itemid, _ := range itembag.GetItems() {
		itemslots := game.Itemslots(itembag.GetItemslots(itemid))

		for _, itemslot := range itemslots {
			itemslot.AddPart(itemslots)
			itemslot.Parts[part.Imaging] = itemslots[0].Parts[part.Imaging]
		}
	}

	player.AddPart(itembag)
}
