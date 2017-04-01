package run

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
	"ospokemon.com/query"
)

func init() {
	event.On(event.PlayersSelect, PlayersSelectMenuBindings)
}

func PlayersSelectMenuBindings(args ...interface{}) {
	player := args[0].(*ospokemon.Player)

	mquery, err := query.BindingsMenusPlayersSelect(player)

	if err != nil {
		log.Add("Username", player.Username).Add("Error", err.Error()).Error("players select menubindings")
		return
	}

	bindings := player.GetBindings()

	if mquery != nil {
		for key, menu := range mquery {
			binding := ospokemon.MakeBinding()
			binding.Key = key

			binding.AddPart(ospokemon.Menu(menu))

			imaging := ospokemon.MakeImaging()
			imaging.Image = "/img/ui/menu/" + menu + ".png"
			binding.AddPart(imaging)

			bindings[key] = binding
		}
	}
}
