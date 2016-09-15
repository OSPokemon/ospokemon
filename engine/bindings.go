package engine

import (
	"github.com/ospokemon/ospokemon/util"
	"time"
)

const COMP_Bindings = "engine/Bindings"

type Binding struct {
	*Action
	Timer *time.Duration
}

type Bindings map[string]*Binding

func (b Bindings) Id() string {
	return COMP_Bindings
}

func (b Binding) Update(u *Universe, e *Entity, d time.Duration) {
	if b.Action.Timer != nil {
		b.Timer = nil
		return
	}

	if *b.Timer < d {
		b.Timer = nil
		util.Event.Fire(EVNT_ActionCast, u, e, b.Action)
	} else {
		*b.Timer -= d
	}
}

func (b Bindings) Update(u *Universe, e *Entity, d time.Duration) {
	for _, binding := range b {
		binding.Update(u, e, d)
	}
}
