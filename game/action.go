package game

import (
	"ospokemon.com/event"
	"ospokemon.com/json"
	"time"
)

const PARTaction = "action"

type Action struct {
	Spell *Spell
	Timer *time.Duration
	Parts
}

func MakeAction() *Action {
	return &Action{
		Parts: make(Parts),
	}
}

func BuildAction(spell *Spell) *Action {
	action := MakeAction()
	action.Spell = spell
	action.AddPart(BuildImaging(spell.Animations))
	return action
}

func (a *Action) Part() string {
	return PARTaction
}

func (parts Parts) GetAction() *Action {
	action, _ := parts[PARTaction].(*Action)
	return action
}

func (a *Action) Update(universe *Universe, entity *Entity, d time.Duration) {
	if a.Timer == nil {
		event.Fire(event.ActionCast, universe, entity, a)
		entity.RemovePart(a)
	}
}

func (a *Action) Json() json.Json {
	return json.Json{
		"timer": json.FmtDuration(a.Timer),
		"spell": a.Spell.Json(),
	}
}
