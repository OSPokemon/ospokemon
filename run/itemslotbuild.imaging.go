package run

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.ItemslotBuild, ItemslotBuildImaging)
}

func ItemslotBuildImaging(args ...interface{}) {
	itemslot := args[0].(*game.Itemslot)
	item, err := query.GetItem(itemslot.Item)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Item":  itemslot.Item,
			"Error": err.Error(),
		}).Error("itemslot build imaging")
		return
	}

	imaging := game.MakeImaging()
	imaging.ReadAnimations(item.Animations)

	itemslot.AddPart(imaging)
}
