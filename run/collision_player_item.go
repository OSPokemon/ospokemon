package run

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/part"
	"github.com/ospokemon/ospokemon/save"
)

func init() {
	event.On(event.MovementCollision, func(args ...interface{}) {
		entity1 := args[0].(*save.Entity)
		entity2 := args[1].(*save.Entity)
		u := args[2].(*save.Universe)

		bag, ok := entity1.Parts[part.ITEMBAG].(*save.Itembag)
		if !ok {
			return
		}

		addslot, ok := entity2.Parts[part.ITEMSLOT].(*save.Itemslot)
		if !ok {
			return
		}

		if bag.Add(addslot) {
			u.Remove(entity2)
		}
	})
}
