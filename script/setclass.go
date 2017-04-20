package script

import (
	"errors"
	"ospokemon.com"
	"ospokemon.com/log"
	"strconv"
)

func init() {
	ospokemon.Scripts["set-class"] = SetClass
	ospokemon.Scripts["setclass"] = SetClass
}

func SetClass(e *ospokemon.Entity, data map[string]interface{}) error {
	var class *ospokemon.Class
	var err error

	switch data_class := data["class"].(type) {
	case string:
		classI, e := strconv.Atoi(data_class)
		if err != nil {
			err = e
			break
		}
		class, err = ospokemon.GetClass(uint(classI))
	case int:
		class, err = ospokemon.GetClass(uint(data_class))
	case *ospokemon.Class:
		class = data_class
	}

	if class == nil {
		return errors.New("set-class: class format")
	} else if err != nil {
		return errors.New("set-class: " + err.Error())
	}

	e.SetClass(class)
	e.GetPlayer().Class = class.Id
	log.Add("EntityId", e.Id).Add("Universe", e.UniverseId).Add("Class", class.Id).Info("setclass")
	return nil
}
