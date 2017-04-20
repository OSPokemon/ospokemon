package run

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
)

func init() {
	event.On(event.BindingDown, BindingDownItemslot)
}

func BindingDownItemslot(args ...interface{}) {
	player := args[0].(*ospokemon.Player)
	binding := args[1].(*ospokemon.Binding)

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

	log.Add("Username", player.Username).Add("Key", binding.Key).Add("Item", item.Id).Info("binding down itemslot")
}
