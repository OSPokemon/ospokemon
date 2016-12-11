package run

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/save"
)

func init() {
	event.On(event.MovementCollision, func(args ...interface{}) {
		entity1 := args[0].(*save.Entity)
		entity2 := args[1].(*save.Entity)
		u := args[2].(*save.Universe)

		bag, ok := entity1.Component(save.COMP_Bag).(*save.Bag)
		if !ok {
			return
		}

		addslot, ok := entity2.Component(save.COMP_ItemSlot).(*save.ItemSlot)
		if !ok {
			return
		}

		if bag.Add(addslot) {
			u.Remove(entity2)
		}
	})
}
