package run

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
)

func init() {
	event.On(event.PlayersSelect, PlayersSelectMenus)
}

func PlayersSelectMenus(args ...interface{}) {
	player := args[0].(*game.Player)
	menus := game.MakeMenus()
	player.AddPart(menus)
}
