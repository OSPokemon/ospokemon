package persistence

import (
	"ospokemon.com"
)

func BindingsPlayersSelect(player *ospokemon.Player) error {
	bindings := player.GetBindings()

	// movement bindings
	mquery, err := BindingsMovementsPlayersSelect(player)
	if err != nil {
		return err
	}
	for key, direction := range mquery {
		binding := ospokemon.MakeBinding()
		binding.Key = key

		binding.AddPart(ospokemon.Walk(direction))

		imaging := ospokemon.MakeImaging()
		imaging.Image = "/img/ui/walk/" + direction + ".png"
		binding.AddPart(imaging)

		bindings[key] = binding
	}

	// menu bindings
	mquery, err = BindingsMenusPlayersSelect(player)
	if err != nil {
		return err
	}
	for key, menu := range mquery {
		binding := ospokemon.MakeBinding()
		binding.Key = key

		binding.AddPart(ospokemon.Menu(menu))

		imaging := ospokemon.MakeImaging()
		imaging.Image = "/img/ui/menu/" + menu + ".png"
		binding.AddPart(imaging)

		bindings[key] = binding
	}

	// itembag bindings
	iquery, err := BindingsItemsPlayersSelect(player)
	if err != nil {
		return err
	}
	itembag := player.GetItembag()
	for key, itemslotid := range iquery {
		binding := ospokemon.MakeBinding()
		binding.Key = key
		itemslot := itembag.Slots[itemslotid]
		itemslot.AddPart(binding)
		binding.Parts = itemslot.Parts

		bindings[key] = binding
	}

	// action bindings
	aquery, err := ActionsBindingsPlayersSelect(player)
	if err != nil {
		return err
	}
	actions := player.GetActions()
	for key, actionid := range aquery {
		binding := ospokemon.MakeBinding()
		binding.Key = key
		action := actions[actionid]

		binding.AddPart(action)
		binding.AddPart(action.GetImaging())

		bindings[key] = binding
	}

	return nil
}

func BindingsPlayersInsert(player *ospokemon.Player) error {
	err := BindingsMovementsPlayersInsert(player)
	if err != nil {
		return err
	}

	err = BindingsMenusPlayersInsert(player)
	if err != nil {
		return err
	}

	err = BindingsItemsPlayersInsert(player)
	if err != nil {
		return err
	}

	err = ActionsBindingsPlayersInsert(player)
	if err != nil {
		return err
	}

	return nil
}

func BindingsPlayersDelete(player *ospokemon.Player) error {
	err := BindingsMovementsPlayersDelete(player)
	if err != nil {
		return err
	}

	err = BindingsMenusPlayersDelete(player)
	if err != nil {
		return err
	}

	err = ActionsBindingsPlayersDelete(player)
	if err != nil {
		return err
	}

	err = BindingsItemsPlayersDelete(player)
	if err != nil {
		return err
	}

	return nil
}
