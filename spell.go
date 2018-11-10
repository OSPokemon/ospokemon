package ospokemon

import (
	"time"

	"ztaylor.me/js"
)

type Spell struct {
	Id         uint
	CastTime   time.Duration
	Cooldown   time.Duration
	Animations map[string]string
	Scripter
}

var spells = make(map[uint]*Spell)

func MakeSpell() *Spell {
	return &Spell{
		Animations: make(map[string]string),
		Scripter:   *MakeScripter(),
	}
}

func (s *Spell) Json() js.Object {
	return map[string]interface{}{
		"id":         s.Id,
		"casttime":   s.CastTime,
		"cooldown":   s.Cooldown,
		"animations": s.Animations,
	}
}

func GetSpell(id uint) (*Spell, error) {
	if spells[id] == nil {
		if s, err := Spells.Select(id); s != nil {
			spells[id] = s
		} else {
			return nil, err
		}
	}

	return spells[id], nil
}

// persistence headers
var Spells struct {
	Select func(uint) (*Spell, error)
}
