package script

import (
	"errors"
	"ospokemon.com"
	"ospokemon.com/script/util"
)

func init() {
	ospokemon.Scripts["itemcast"] = ItemCast
}

func ItemCast(e *ospokemon.Entity, data map[string]interface{}) error {
	itembag := e.GetItembag()
	if itembag == nil {
		return errors.New("itemcast: itembag missing")
	}

	item, err := util.GetItem(data["item"])
	if err != nil {
		return errors.New("itemcast: " + err.Error())
	}

	if itembag.Timers[item.Id] != nil {
		return errors.New("itemcast: item cooldown")
	}

	itemslot := itembag.Slots[item.Id]
	if itemslot == nil {
		return errors.New("itemcast: itemslot missing")
	}

	timer := item.CastTime + item.Cooldown
	itembag.Timers[itemslot.Item.Id] = &timer
	return nil
}
