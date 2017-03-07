package game

import (
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

	return a
}

func BuildAction(spell *Spell) *Action {
	action := MakeAction()

	imaging := MakeImaging()
	imaging.ReadAnimations(spell.Animations)
	action.AddPart(imaging)

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
