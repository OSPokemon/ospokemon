package engine

import (
	"github.com/ospokemon/ospokemon/util"
	"time"
)

const EVNT_ActionCast = "ospokemon/engine/Action.Cast"

type Action struct {
	Name     string
	Image    string
	ScriptId string
	CastTime time.Duration
	CoolDown time.Duration
	Timer    *time.Duration
}

func (a *Action) Update(u *Universe, e *Entity, d time.Duration) {
	if a.Timer == nil {
		return
	}

	*a.Timer += d

	if *a.Timer-d <= a.CastTime && *a.Timer >= a.CastTime {
		a.Cast(u, e)
	}

	if *a.Timer > a.CastTime+a.CoolDown {
		a.Timer = nil
	}
}

func (a *Action) Cast(u *Universe, e *Entity) {
	util.Event.Fire(EVNT_ActionCast, u, e, a)
	u.Fire(EVNT_ActionCast, u, e, a)
	e.Fire(EVNT_ActionCast, u, e, a)
}
