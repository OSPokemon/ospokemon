package game

import (
	"time"
)

type Spell struct {
	Id         uint
	Script     string
	CastTime   time.Duration
	Cooldown   time.Duration
	Animations map[string]string
	Data       map[string]string
}

var Spells = make(map[uint]*Spell)

func MakeSpell(id uint) *Spell {
	s := &Spell{
		Id:         id,
		Animations: make(map[string]string),
		Data:       make(map[string]string),
	}

	return s
}

func (s *Spell) Snapshot() map[string]interface{} {
	return map[string]interface{}{
		"id":         s.Id,
		"casttime":   s.CastTime,
		"cooldown":   s.Cooldown,
		"animations": s.Animations,
	}
}
