package json

import (
	"github.com/ospokemon/ospokemon/game"
	"strconv"
)

func Itembag(b *game.Itembag) (string, map[string]interface{}) {
	data := make(map[string]interface{})

	for pos, slot := range b.Slots {
		key := strconv.Itoa(pos)
		if slot == nil {
			data[key] = nil
			continue
		}

		_, slotData := Itemslot(slot)

		if b.Timers[slot.Item] != nil {
			slotData["timer"] = *b.Timers[slot.Item]
		} else {
			slotData["timer"] = 0
		}

		data[key] = slotData
	}

	return "itembag", data
}
