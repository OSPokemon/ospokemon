package ospokemon

import (
	"time"

	"ztaylor.me/cast"
	"ztaylor.me/js"
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
		if t := action.Timer; t != nil {
		} else if td := t.Duration(); td < d {
			action.Timer = nil
		} else {
			action.Timer.Set(td - d)
		}
	}
}

func (actions Actions) Json() js.Object {
	data := js.Object{}
	for id, action := range actions {
		data[cast.String(id)] = action.Json()
	}
	return data
}
