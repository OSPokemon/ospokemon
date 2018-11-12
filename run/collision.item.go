package run

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/script"
	"ztaylor.me/log"
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
			log.Add("Entity", entity1.Id).Add("Universe", entity1.UniverseId).Add("Error", err.Error()).Error("collision.item")
			return
		}

		universe := ospokemon.Universes.Cache[entity2.UniverseId]
		universe.Remove(entity2)
	})
}
