package script

import (
	"errors"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/part"
	"github.com/ospokemon/ospokemon/query"
	"strconv"
)

func init() {
	game.Scripts["itemchange"] = func(e *game.Entity, data map[string]string) error {
		itembag, ok := e.Parts[part.Itembag].(*game.Itembag)
		if !ok {
			return errors.New("itemchange: itembag mising")
		}

		itemid, err := strconv.ParseUint(data["item"], 10, 0)
		if err != nil {
			return err
		}

		amount, err := strconv.Atoi(data["amount"])
		if err != nil {
			return err
		}

		item, err := query.GetItem(uint(itemid))
		if err != nil {
			return err
		}

		return ItemChange(itembag, item, amount)
	}
}

func ItemChange(itembag *game.Itembag, item *game.Item, amount int) error {
	if amount > 0 {
		if !itembag.Add(item, int(amount)) {
			return errors.New("itemchange: itembag full")
		}
	} else if amount < 0 {
		if !itembag.Remove(item, -int(amount)) {
			return errors.New("itemchange: missing items")
		}
	}
	return nil
}
