package ospokemon

import (
	"ospokemon.com/json"
	"time"
)

const PARTactions = "actions"

type Actions map[uint]*Action

func (a Actions) Part() string {
	return PARTactions
}

func (parts Parts) GetActions() Actions {
	actions, _ := parts[PARTactions].(Actions)
	return actions
}

func (actions Actions) Update(universe *Universe, entity *Entity, d time.Duration) {
	for _, action := range actions {
		if action.Timer == nil {
		} else if *action.Timer < d {
			action.Timer = nil
		} else {
			*action.Timer = *action.Timer - d
		}
	}
}

func (actions Actions) Json() json.Json {
	data := json.Json{}
	for id, action := range actions {
		data[json.StringUint(id)] = action.Json()
	}
	return data
}
