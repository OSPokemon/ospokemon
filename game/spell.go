package game

import (
	"time"
)

type Spell struct {
	Id         uint
	CastTime   time.Duration
	Cooldown   time.Duration
	Animations map[string]string
	Scripter
}

var Spells = make(map[uint]*Spell)

func MakeSpell() *Spell {
	return &Spell{
		Animations: make(map[string]string),
		Scripter:   MakeScripter(),
	}
}

func (s *Spell) Snapshot() map[string]interface{} {
	return map[string]interface{}{
		"id":         s.Id,
		"casttime":   s.CastTime,
		"cooldown":   s.Cooldown,
		"animations": s.Animations,
	}
}
