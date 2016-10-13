package save

import (
	"time"
)

type Spell struct {
	Id       uint
	Image    string
	ScriptId uint
	CastTime time.Duration
	Cooldown time.Duration
	Data     map[string]string
}

var Spells = make(map[uint]*Spell)
