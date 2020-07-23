package ospokemon

import (
	"time"

	"taylz.io/types"
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

func (actions Actions) Json() types.Dict {
	data := types.Dict{}
	for id, action := range actions {
		data[types.StringUint(id)] = action.Json()
	}
	return data
}
