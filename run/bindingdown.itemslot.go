package run

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
)

func init() {
	event.On(event.BindingDown, BindingDownItemslot)
}

func BindingDownItemslot(args ...interface{}) {
	player := args[0].(*game.Player)
	binding := args[1].(*game.Binding)

	itemslot := binding.GetItemslot()
	if itemslot == nil {
		return
	}

	itembag := player.GetItembag()

	if itembag.Timers[itemslot.Item.Id] != nil {
		return
	}

	item := itemslot.Item
	timer := item.CastTime + item.Cooldown
	itembag.Timers[itemslot.Item.Id] = &timer

	logrus.WithFields(logrus.Fields{
		"Username": player.Username,
		"Binding":  binding,
	}).Info("binding down itemslot")
}
