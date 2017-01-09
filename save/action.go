package save

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/part"
	"time"
)

type Action struct {
	Spell uint
	Timer *time.Duration
	part.Parts
}

func MakeAction() *Action {
	a := &Action{
		Parts: make(part.Parts),
	}

	a.AddPart(a)

	event.Fire(event.ActionMake, a)

	return a
}

func (a *Action) Part() string {
	return part.ACTION
}

func (a *Action) Cast(b *Binding) {
	if spell, err := GetSpell(a.Spell); spell != nil {
		timer := spell.CastTime + spell.Cooldown
		a.Timer = &timer
	} else if err != nil {
		logrus.Error(err.Error())
	}
}

func (a *Action) Update(u *Universe, e *Entity, d time.Duration) {
	if a.Timer == nil {
		return
	}

	if *a.Timer < d {
		a.Timer = nil
	}
	*a.Timer = *a.Timer - d
}

func (a *Action) Json(expand bool) (string, map[string]interface{}) {
	data := make(map[string]interface{})

	if a.Timer != nil {
		data["timer"] = int64(*a.Timer)
	} else {
		data["timer"] = 0
	}

	if spell, _ := GetSpell(a.Spell); spell != nil {
		data["spell"] = spell.Snapshot()
	} else {
		data["spell"] = a.Spell
	}

	if expand {
		for _, part := range a.Parts {
			if jsoner, ok := part.(Jsoner); ok {
				key, jsonerData := jsoner.Json(false)
				data[key] = jsonerData
			}
		}
	}

	return "action", data
}
