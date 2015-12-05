package world

import (
	"time"
)

// Effects

type EffectScript func(*Effect, Entity, time.Time)

const (
	PRIOstate    = -1
	PRIOstandard = 0
	PRIOafter    = 1
)

type Effect struct {
	Name     string
	Priority int
	Data     interface{}
	Script   EffectScript
	Start    time.Time
	Duration time.Duration
}

// Effects sorting

type Effects []*Effect

func (e Effects) Len() int {
	return len(e)
}

func (e Effects) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e Effects) Less(i, j int) bool {
	return e[i].Priority < e[j].Priority
}

// Spells

type TargetType uint8
type SpellScript func(Entity, interface{}, time.Time)

const (
	TRGTnone TargetType = iota
	TRGTposition
	TRGTentity
)

type Spell struct {
	Id         int
	Name       string
	CastTime   time.Duration
	Cooldown   time.Duration
	MoveCast   bool
	Cost       SpellCost
	Range      float64
	TargetType TargetType
	Graphic    string
	Script     SpellScript
}

func (s *Spell) String() string {
	return s.Name
}

type SpellCost struct {
	Mana  int
	Items map[int]int
}

// Abilities

type Ability struct {
	Spell    *Spell
	LastCast time.Time
	CastTime time.Duration
	Cooldown time.Duration
	MoveCast bool
	Cost     SpellCost
	Range    float64
}

func (a *Ability) String() string {
	return a.Spell.String()
}

// Actions

type Action struct {
	Ability *Ability
	Target  interface{}
}

// Controls

const (
	CTRLimmune uint8 = 0x01
	CTRLstasis uint8 = 0x02
	CTRLstun   uint8 = 0x04
	CTRLroot   uint8 = 0x08
	CTRLcloak  uint8 = 0x10
	// pseudo control states
	CTRLPprotected uint8 = CTRLimmune | CTRLstasis
	CTRLPnocast    uint8 = CTRLstasis | CTRLstun
	CTRLPstuck     uint8 = CTRLstasis | CTRLroot
)

type Controls struct {
	Action    *Action
	State     uint8
	Abilities map[string]*Ability
}
