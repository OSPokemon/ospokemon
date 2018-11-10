package run

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ztaylor.me/log"
	"ospokemon.com/script"
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

		universe := ospokemon.Multiverse[entity2.UniverseId]
		universe.Remove(entity2)
	})
}
