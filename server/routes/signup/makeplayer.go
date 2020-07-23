package signup

import (
	"github.com/ospokemon/ospokemon"
)

func MakePlayer(account *ospokemon.Account) {
	class, err := ospokemon.GetClass(0)
	if err != nil {
		ospokemon.LOG().Add("Error", err.Error()).Error("signup: failed to get class 0")
		return
	}
	entity := ospokemon.MakeEntity()
	player := ospokemon.BuildPlayer(account.Username, class, entity)
	player.Username = account.Username

	bindings := player.GetBindings()

	for key, menu := range menuBindings {
		binding := ospokemon.MakeBinding()
		binding.Key = key
		binding.AddPart(menu)
		bindings[binding.Key] = binding
	}

	for key, direction := range movementBindings {
		binding := ospokemon.MakeBinding()
		binding.Key = key
		binding.AddPart(direction)
		bindings[binding.Key] = binding
	}

	stats := player.GetStats()
	stats["speed"] = &ospokemon.Stat{Value: 5, Base: 5}

	player.AddPart(account)
	player.AddPart(player)
	account.Parts = player.Parts
}
