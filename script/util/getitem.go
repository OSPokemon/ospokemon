package util

import (
	"errors"
	"ospokemon.com"
	"strconv"
)

func GetItem(data interface{}) (*ospokemon.Item, error) {
	var item *ospokemon.Item
	var err error

	switch data_item := data.(type) {
	case nil:
		err = errors.New("getitem: nil")
	case *ospokemon.Item:
		item = data_item
	case int:
		item, err = ospokemon.GetItem(uint(data_item))
	case uint:
		item, err = ospokemon.GetItem(data_item)
	case float64:
		item, err = ospokemon.GetItem(uint(data_item))
	case string:
		itemid64, err2 := strconv.ParseUint(data_item, 10, 0)
		if err2 == nil {
			item, err = ospokemon.GetItem(uint(itemid64))
		} else {
			err = errors.New("getitem: " + err2.Error())
		}
	default:
		err = errors.New("getitem: \"item\" type unrecognized")
	}

	return item, err
}
