package run

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
	"ospokemon.com/persistence"
)

func init() {
	event.On(event.EntitiesUniversesSelect, EntitiesUniversesSelectDialogs)
}

func EntitiesUniversesSelectDialogs(args ...interface{}) {
	entities := args[0].(map[uint]*ospokemon.Entity)
	universe := args[1].(*ospokemon.Universe)
	dialogs, err := persistence.DialogsSelect(universe)

	if err != nil {
		log.Add("Universe", universe.Id).Add("Error", err.Error()).Error("entities universes select dialogs")
		return
	}

	for entityId, dialog := range dialogs {
		entities[entityId].AddPart(dialog)
	}
}
