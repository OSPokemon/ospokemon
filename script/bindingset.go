package script

import (
	"errors"
	"ospokemon.com"
	"strconv"
)

func init() {
	ospokemon.Scripts["bindingset"] = BindingSet
}

func BindingSet(e *ospokemon.Entity, data map[string]interface{}) error {
	bindings := e.GetBindings()
	if bindings == nil {
		return errors.New("bindingset: bindings missing")
	}

	key, ok := data["key"].(string)
	if !ok {
		return errors.New("bindingset: key missing")
	}

	var spellid uint
	switch data_spell := data["spell"].(type) {
	case int:
		spellid = uint(data_spell)
		break
	case uint:
		spellid = data_spell
		break
	case nil:
		break
	case string:
		spellid64, err := strconv.ParseUint(data_spell, 10, 0)
		if err == nil {
			spellid = uint(spellid64)
			break
		}
	default:
		return errors.New("bindingset: spell format")
	}

	if spellid > 0 {
		if actions := e.GetActions(); actions != nil {
			if action := actions[spellid]; action != nil {
				bindings.SetAction(key, action)
				return nil
			}
			return errors.New("bindingset: action missing")
		}
		return errors.New("bindingset: actions missing")
	}

	var itemid uint
	switch data_itemid := data["itemid"].(type) {
	case int:
		itemid = uint(data_itemid)
		break
	case uint:
		itemid = data_itemid
		break
	case nil:
		break
	case string:
		itemidi, err := strconv.Atoi(data_itemid)
		if err != nil {
			itemid = uint(itemidi)
			break
		}
	case float64:
		itemid = uint(data_itemid)
		break
	default:
		return errors.New("bindingset: itemslot format")
	}

	if !(itemid < 0) {
		if itembag := e.GetItembag(); itembag != nil {
			if itemslot := itembag.Slots[itemid]; itemslot != nil {
				bindings.SetItemslot(key, itemslot)
				return nil
			}
			return errors.New("bindingset: itemslot missing")
		}
		return errors.New("bindingset: itembag missing")
	}

	if walk, _ := data["walk"].(string); walk != "" {
		if movement := e.GetMovement(); movement != nil {
			bindings.SetWalk(key, walk)
			return nil
		}
		return errors.New("bindingset: movement missing")
	}

	if menu, _ := data["menu"].(string); menu != "" {
		if menus := e.GetMenus(); menus != nil {
			bindings.SetMenu(key, menu)
			return nil
		}
		return errors.New("bindingset: menus missing")
	}

	return errors.New("bindingset: unrecognized")
}
