package script

import (
	"errors"
	"ospokemon.com"
	"ospokemon.com/script/util"
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

	item, err := util.GetItem(data["item"])
	if err != nil {
		return errors.New("itemchange: " + err.Error())
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
