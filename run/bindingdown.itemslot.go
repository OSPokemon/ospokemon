package run

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/part"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.BindingDown, BindingDownItemslot)
}

func BindingDownItemslot(args ...interface{}) {
	player := args[0].(*game.Player)
	binding := args[1].(*game.Binding)

	if itemslot, ok := binding.Parts[part.Itemslot].(*game.Itemslot); ok {
		itembag := player.Parts[part.Itembag].(*game.Itembag)
		if itembag.Timers[itemslot.Item] != nil {
			return
		} else if item, err := query.GetItem(itemslot.Item); item != nil {
			timer := item.CastTime + item.Cooldown
			itembag.Timers[itemslot.Item] = &timer

			logrus.WithFields(logrus.Fields{
				"Username": player.Username,
				"Binding":  binding,
			}).Info("binding down itemslot")
		} else if err != nil {
			logrus.WithFields(logrus.Fields{
				"Username": player.Username,
				"Binding":  binding,
				"Error":    err.Error(),
			}).Error("binding down itemslot")
		}
	}
}
