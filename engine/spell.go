package engine

import (
	"time"
)

var Spells = make(map[int]*Spell)
var LoadSpell func(spellId int) (*Spell, error)

type Spell struct {
	Id          int
	Name        string
	CastTime    time.Duration
	MoveCast    bool
	ChannelTime time.Duration
	MoveChannel bool
	Cooldown    time.Duration
	ManaCost    int
	ItemCost    map[int]int
	TargetType  string
	Range       int
	Size        int
	Graphic     string
	Script      SpellScript
}

func GetSpell(spellId int) (*Spell, error) {
	if Spells[spellId] != nil {
		return Spells[spellId], nil
	}

	spell, err := LoadSpell(spellId)

	if err == nil {
		Spells[spellId] = spell
	}

	return spell, err
}
