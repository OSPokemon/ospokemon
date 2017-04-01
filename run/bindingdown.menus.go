package run

import (
	"ospokemon.com/event"
	"ospokemon.com/game"
)

func init() {
	event.On(event.BindingDown, BindingDownMenus)
}

func BindingDownMenus(args ...interface{}) {
	player := args[0].(*game.Player)
	binding := args[1].(*game.Binding)
	menus := player.GetMenus()

	if menu := binding.GetMenu(); menu != "" {
		menus.Toggle(menu)
	}
}
