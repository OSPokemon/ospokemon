package run

import (
	"ospokemon.com/event"
	"ospokemon.com/game"
	"ospokemon.com/log"
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

	log.Add("Username", player.Username).Add("Binding", binding).Info("binding down itemslot")
}
