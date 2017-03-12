package script

import (
	"errors"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/query"
	"strconv"
)

func init() {
	game.Scripts["itemchange"] = ItemChange
}

func ItemChange(e *game.Entity, data map[string]interface{}) error {
	itembag := e.GetItembag()
	if itembag == nil {
		return errors.New("itemchange: itembag mising")
	}

	var item *game.Item
	var err error

	switch data_item := data["item"].(type) {
	case *game.Item:
		item = data_item
		break
	case uint:
		item, err = query.GetItem(data_item)
		break
	case string:
		itemid64, err := strconv.ParseUint(data_item, 10, 0)
		if err == nil {
			item, err = query.GetItem(uint(itemid64))
		}
		break
	default:
		err = errors.New("itemchange: item format")
	}

	if err != nil {
		return err
	}

	var amount int
	switch data_amount := data["amount"].(type) {
	case int:
		amount = data_amount
		break
	case string:
		amounti, err := strconv.Atoi(data_amount)
		if err == nil {
			amount = amounti
			break
		}
	default:
		return errors.New("itemchange: amount format")
	}

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
