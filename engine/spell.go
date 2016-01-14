package engine

import (
	log "github.com/Sirupsen/logrus"
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

func GetSpell(spellId int) *Spell {
	if Spells[spellId] == nil {
		if spell, err := LoadSpell(spellId); err == nil {
			Spells[spellId] = spell
		} else {
			log.WithFields(log.Fields{
				"SpellId": spellId,
				"Error":   err.Error(),
			}).Info("Spell lookup failed")
		}
	}

	return Spells[spellId]
}
