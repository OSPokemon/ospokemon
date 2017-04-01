package run

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
	"ospokemon.com/query"
)

func init() {
	event.On(event.PlayersSelect, PlayersSelectMovementBindings)
}

func PlayersSelectMovementBindings(args ...interface{}) {
	player := args[0].(*ospokemon.Player)

	mquery, err := query.BindingsMovementsPlayersSelect(player)

	if err != nil {
		log.Add("Username", player.Username).Add("Error", err.Error()).Error("players select movement bindings")
		return
	}

	bindings := player.GetBindings()

	if mquery != nil {
		for key, direction := range mquery {
			binding := ospokemon.MakeBinding()
			binding.Key = key

			binding.AddPart(ospokemon.Walk(direction))

			imaging := ospokemon.MakeImaging()
			imaging.Image = "/img/ui/walk/" + direction + ".png"
			binding.AddPart(imaging)

			bindings[key] = binding
		}
	}
}
