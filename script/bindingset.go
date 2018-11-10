package script

import (
	"errors"
	"ospokemon.com"
	"ztaylor.me/log"
	"ospokemon.com/script/util"
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

	// spell binding
	if data["spell"] != nil {
		spell, err := util.GetSpell(data["item"])
		if err != nil {
			return errors.New("bindingset: " + err.Error())
		}

		if actions := e.GetActions(); actions == nil {
			return errors.New("bindingset: actions missing")
		} else if action := actions[spell.Id]; action == nil {
			return errors.New("bindingset: action missing")
		} else {
			log.Add("Username", e.GetPlayer().Username).Add("Key", key).Add("Spell", action.Spell.Id).Info("bindingset: spell")
			bindings.SetAction(key, action)
			return nil
		}
	}

	// item binding
	if data["item"] != nil {
		item, err := util.GetItem(data["item"])
		if err != nil {
			return errors.New("bindingset: " + err.Error())
		}

		if itembag := e.GetItembag(); itembag == nil {
			return errors.New("bindingset: itembag missing")
		} else if itemslot := itembag.Slots[item.Id]; itemslot == nil {
			return errors.New("bindingset: itemslot missing")
		} else {
			log.Add("Username", e.GetPlayer().Username).Add("Key", key).Add("Item", itemslot.Item.Id).Info("bindingset: itemslot")
			bindings.SetItemslot(key, itemslot)
			return nil
		}
	}

	// walk binding
	if walk, _ := data["walk"].(string); walk != "" {
		if movement := e.GetMovement(); movement != nil {
			bindings.SetWalk(key, walk)
			return nil
		}
		return errors.New("bindingset: movement missing")
	}

	// menu binding
	if menu, _ := data["menu"].(string); menu != "" {
		if menus := e.GetMenus(); menus != nil {
			bindings.SetMenu(key, menu)
			return nil
		}
		return errors.New("bindingset: menus missing")
	}

	return errors.New("bindingset: unrecognized")
}
