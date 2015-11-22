package world

import (
	"time"
)

// Effects

type EffectType uint8

const (
	EFCTimmune EffectType = iota
	EFCTstasis
	EFCThealth
	EFCTstun
	EFCTroot
	EFCTmove
)

type Effect struct {
	Name     string
	Type     EffectType
	Data     interface{}
	Start    time.Time
	Duration time.Duration
}

type Effects []*Effect

func (e Effects) Len() int {
	return len(e)
}

func (e Effects) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e Effects) Less(i, j int) bool {
	return e[i].Type < e[j].Type
}

// Spells

type TargetType uint8
type SpellScript func(Entity, interface{})

const (
	TRGTnone TargetType = iota
	TRGTposition
	TRGTentity
)

type Spell struct {
	Name        string
	Description string
	CastTime    time.Duration
	Cooldown    time.Duration
	Cost        *SpellCost
	Range       float64
	TargetType  TargetType
	Script      SpellScript
}

type SpellCost struct {
	Mana  int
	Items map[int]int
}

// Abilities

type Ability struct {
	LastCast time.Time
	Spell    *Spell
}

// Actions

type Action struct {
	Clock   time.Time
	Target  interface{}
	Ability *Ability
}

// Controls

const (
	CTRLimmune uint8 = 0x01
	CTRLstasis uint8 = 0x02
	CTRLstun   uint8 = 0x04
	CTRLroot   uint8 = 0x08
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

var WalkAbility = &Ability{
	Spell: &Spell{"walk", "", 0, 0, nil, 0, TRGTnone, nil},
} // Special flag ability
