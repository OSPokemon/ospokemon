package script

import (
	"errors"
	"ospokemon.com"
	"strconv"
)

func init() {
	ospokemon.Scripts["itemchange"] = ItemChange
}

func ItemChange(e *ospokemon.Entity, data map[string]interface{}) error {
	itembag := e.GetItembag()
	if itembag == nil {
		return errors.New("itemchange: itembag mising")
	}

	toaster := e.GetToaster()
	if toaster == nil {
		return errors.New("itemchange: toaster missing")
	}

	var item *ospokemon.Item
	var err error

	switch data_item := data["item"].(type) {
	case *ospokemon.Item:
		item = data_item
		break
	case uint:
		item, err = ospokemon.GetItem(data_item)
		break
	case string:
		itemid64, err := strconv.ParseUint(data_item, 10, 0)
		if err == nil {
			item, err = ospokemon.GetItem(uint(itemid64))
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
		if itembag.Add(item, int(amount)) {
			toaster.Add(&ospokemon.Toast{
				Color:   "green",
				Image:   item.Animations["portrait"],
				Message: "+" + strconv.Itoa(amount),
			})
		} else {
			return errors.New("itemchange: itembag full")
		}
	} else if amount < 0 {
		if itembag.Remove(item, -int(amount)) {
			toaster.Add(&ospokemon.Toast{
				Color:   "orange",
				Image:   item.Animations["portrait"],
				Message: strconv.Itoa(amount),
			})
		} else {
			return errors.New("itemchange: missing items")
		}
	}
	return nil
}
