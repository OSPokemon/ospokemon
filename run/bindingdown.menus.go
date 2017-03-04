package run

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/part"
)

func init() {
	event.On(event.BindingDown, BindingDownMenus)
}

func BindingDownMenus(args ...interface{}) {
	player := args[0].(*game.Player)
	binding := args[1].(*game.Binding)
	menus := player.Parts[part.Menus].(game.Menus)

	if menu, ok := binding.Parts[part.Menu].(game.Menu); ok {
		menus.Toggle(menu)
	}
}
