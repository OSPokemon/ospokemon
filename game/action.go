package game

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/json"
	"github.com/ospokemon/ospokemon/part"
	"time"
)

type Action struct {
	Spell *Spell
	Timer *time.Duration
	part.Parts
}

func MakeAction() *Action {
	return &Action{
		Parts: make(part.Parts),
	}
}

func BuildAction(spell *Spell) *Action {
	action := MakeAction()
	action.Spell = spell
	action.AddPart(BuildImaging(spell.Animations))
	return action
}

func (a *Action) Part() string {
	return part.Action
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
