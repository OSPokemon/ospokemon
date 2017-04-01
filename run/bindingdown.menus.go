package run

import (
	"ospokemon.com"
	"ospokemon.com/event"
)

func init() {
	event.On(event.BindingDown, BindingDownMenus)
}

func BindingDownMenus(args ...interface{}) {
	player := args[0].(*ospokemon.Player)
	binding := args[1].(*ospokemon.Binding)
	menus := player.GetMenus()

	if menu := binding.GetMenu(); menu != "" {
		menus.Toggle(menu)
	}
}
