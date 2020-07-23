package run

import (
	"github.com/ospokemon/ospokemon"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/script"
)

func init() {
	event.On(event.Collision, func(args ...interface{}) {
		entity1 := args[0].(*ospokemon.Entity)
		entity2 := args[1].(*ospokemon.Entity)

		itemslot := entity2.GetItemslot()
		if itemslot == nil {
			return
		}

		err := script.ItemChange(entity1, map[string]interface{}{
			"item":   itemslot.Item,
			"amount": itemslot.Amount,
		})

		if err != nil {
			ospokemon.LOG().Add("Entity", entity1.Id).Add("Universe", entity1.UniverseId).Add("Error", err.Error()).Error("collision.item")
			return
		}

		universe := ospokemon.Universes.Cache[entity2.UniverseId]
		universe.Remove(entity2)
	})
}
