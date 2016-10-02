package engine

import (
	"time"
)

const COMP_Spells = "engine/Spells"
const EVNT_SpellPull = "engine.Spell.Pull"

type Spell struct {
	Id       uint
	Image    string
	ScriptId uint
	CastTime time.Duration
	Cooldown time.Duration
	Data     map[string]string
}

type Spells map[uint]*Spell
