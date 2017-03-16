package run

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
	"github.com/ospokemon/ospokemon/script"
)

func init() {
	event.On(event.Collision, func(args ...interface{}) {
		entity1 := args[0].(*game.Entity)
		entity2 := args[1].(*game.Entity)

		itemslot := entity2.GetItemslot()
		if itemslot == nil {
			return
		}

		err := script.ItemChange(entity1, map[string]interface{}{
			"item":   itemslot.Item,
			"amount": itemslot.Amount,
		})

		if err != nil {
			log.Add("Entity", entity1.Id).Add("Universe", entity1.UniverseId).Add("Error", err.Error()).Error("collision.item")
			return
		}

		universe := game.Multiverse[entity2.UniverseId]
		universe.Remove(entity2)
	})
}
