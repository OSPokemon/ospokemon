package script

import (
	"errors"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/part"
	"strconv"
)

func init() {
	game.Scripts["bindingset"] = BindingSet
}

func BindingSet(e *game.Entity, data map[string]interface{}) error {
	bindings, ok := e.Parts[part.Bindings].(game.Bindings)
	if !ok {
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
		if actions, _ := e.Parts[part.Actions].(game.Actions); actions != nil {
			if action := actions[spellid]; action != nil {
				bindings.SetAction(key, action)
				return nil
			}
			return errors.New("bindingset: action missing")
		}
		return errors.New("bindingset: actions missing")
	}

	itemslotid := -1
	switch data_itemslot := data["itemslot"].(type) {
	case int:
		itemslotid = data_itemslot
		break
	case nil:
		break
	case string:
		itemslotidi, err := strconv.Atoi(data_itemslot)
		if err != nil {
			itemslotid = itemslotidi
			break
		}
	case float64:
		itemslotid = int(data_itemslot)
		break
	default:
		return errors.New("bindingset: itemslot format")
	}

	if !(itemslotid < 0) {
		if itembag, _ := e.Parts[part.Itembag].(*game.Itembag); itembag != nil {
			if itemslot := itembag.Slots[itemslotid]; itemslot != nil {
				bindings.SetItemslot(key, itemslot)
				return nil
			}
			return errors.New("bindingset: itemslot missing")
		}
		return errors.New("bindingset: itembag missing")
	}

	if walk, _ := data["walk"].(string); walk != "" {
		if movement, _ := e.Parts[part.Movement].(*game.Movement); movement != nil {
			bindings.SetWalk(key, walk)
			return nil
		}
		return errors.New("bindingset: movement missing")
	}

	if menu, _ := data["menu"].(string); menu != "" {
		if menus, _ := e.Parts[part.Menus].(*game.Menus); menus != nil {
			bindings.SetMenu(key, menu)
			return nil
		}
		return errors.New("bindingset: menus missing")
	}

	return errors.New("bindingset: unrecognized")
}
