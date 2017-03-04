package script

import (
	"errors"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/part"
	"github.com/ospokemon/ospokemon/query"
	"strconv"
)

const ITEMCHANGE = "itemchange"

func init() {
	game.Scripts[ITEMCHANGE] = func(e *game.Entity, data map[string]string) error {
		itembag, ok := e.Parts[part.Itembag].(*game.Itembag)
		if !ok {
			return errors.New("itembag mising")
		}

		itemid, err := strconv.ParseUint(data["item"], 10, 64)
		if err != nil {
			return err
		}

		amount, err := strconv.ParseInt(data["amount"], 10, 64)
		if err != nil {
			return err
		}

		item, err := query.GetItem(uint(itemid))
		if err != nil {
			return err
		}

		return ItemChange(itembag, item, int(amount))
	}
}

func ItemChange(itembag *game.Itembag, item *game.Item, amount int) error {
	if amount > 0 {
		return itemGive(itembag, item, amount)
	} else if amount < 0 {
		return itemTake(itembag, item, amount)
	}
	return errors.New("script.ItemChange amount is 0")
}

func itemGive(itembag *game.Itembag, item *game.Item, amount int) error {
	for _, itemslot := range itembag.GetItemslots(item.Id) {
		itemslot.Amount += amount

		if itemslot.Amount < item.Stack {
			return nil
		}
		amount = itemslot.Amount - item.Stack
	}

	for id, itemslot := range itembag.Slots {
		if itemslot != nil {
			continue
		}

		itemslot = game.MakeItemslot()
		itemslot.Id = id
		itemslot.Item = item.Id
		itemslot.Amount = amount
		amount = 0

		imaging := game.MakeImaging()
		imaging.ReadAnimations(item.Animations)
		itemslot.AddPart(imaging)
		itembag.Slots[id] = itemslot
		return nil
	}

	return errors.New("script.ItemChange bag is full!")
}

func itemTake(itembag *game.Itembag, item *game.Item, amount int) error {
	if itembag.GetItems()[item.Id] < -amount {
		return errors.New("script.ItemChange items missing")
	}

	for _, itemslot := range itembag.GetItemslots(item.Id) {
		itemslot.Amount += amount

		if itemslot.Amount < 0 {
			amount = itemslot.Amount
			itemslot.Amount = 0
		} else {
			if itemslot.Amount == 0 {
				itembag.Slots[itemslot.Id] = nil
			}
			return nil
		}
	}

	return errors.New("script.ItemChange items missing!")
}
