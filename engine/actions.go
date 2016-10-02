package engine

import (
	"time"
)

const COMP_Actions = "engine/Actions"
const EVNT_ActionCast = "engine.Action.Cast"

type Action struct {
	SpellId uint
	Timer   *time.Duration
}

type Actions map[uint]*Action

func MakeAction(spellId uint) *Action {
	return &Action{
		SpellId: spellId,
	}
}

func (a Actions) Id() string {
	return COMP_Actions
}

func (a Action) Update(u *Universe, e *Entity, d time.Duration) {
	if a.Timer != nil {
		if *a.Timer < d {
			a.Timer = nil
		} else {
			*a.Timer -= d
		}
	}
}

func (a Actions) Update(u *Universe, e *Entity, d time.Duration) {
	for _, action := range a {
		action.Update(u, e, d)
	}
}

func (a Action) Snapshot() map[string]interface{} {
	timebuff := 0
	if a.Timer != nil {
		timebuff = int(*a.Timer)
	}

	return map[string]interface{}{
		"spellid": a.SpellId,
		"timer":   timebuff,
	}
}

func (a Actions) Snapshot() map[string]interface{} {
	return nil
}
