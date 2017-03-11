package game

import (
	"github.com/ospokemon/ospokemon/json"
	"github.com/ospokemon/ospokemon/part"
	"time"
)

type Actions map[uint]*Action

func (a Actions) Part() string {
	return part.Actions
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
