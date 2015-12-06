package world

import (
	"time"
)

// Effects

type TargetType uint8
type SpellScript func(Entity, interface{}, time.Time)
type EffectScript func(*Effect, Entity, time.Time)

const (
	CTRLdead   uint8 = 0x01
	CTRLimmune uint8 = 0x02
	CTRLstasis uint8 = 0x04
	CTRLstun   uint8 = 0x08
	CTRLroot   uint8 = 0x10
	CTRLcloak  uint8 = 0x20
	// pseudo control states
	CTRLPprotected uint8 = CTRLimmune | CTRLstasis | CTRLdead
	CTRLPnocast    uint8 = CTRLstasis | CTRLstun | CTRLdead
	CTRLPstuck     uint8 = CTRLstasis | CTRLroot | CTRLdead
)
const (
	TRGTnone TargetType = iota
	TRGTentity
	TRGTprojectile
)
const (
	PRIOstate    = -1
	PRIOstandard = 0
	PRIOafter    = 1
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
	TargetData map[string]string
	Graphic    string
	Script     SpellScript
}

type SpellCost struct {
	Mana  int
	Items map[int]int
}

type Ability struct {
	Spell      *Spell
	LastCast   time.Time
	CastTime   time.Duration
	Cooldown   time.Duration
	MoveCast   bool
	Cost       SpellCost
	Range      float64
	TargetData map[string]string
}

type Effect struct {
	Name     string
	Priority int
	Data     map[string]string
	Script   EffectScript
	Start    *time.Time
	Duration time.Duration
}

type Action struct {
	Ability *Ability
	Start   *time.Time
	Target  interface{}
}

// Useful controls methods

func IsDead(m Mortality) bool {
	return m.Control()&CTRLdead > 0
}

func IsProtected(m Mortality) bool {
	return m.Control()&CTRLPprotected > 0
}

func NoCast(m Mortality) bool {
	return m.Control()&CTRLPnocast > 0
}

func IsStuck(m Mortality) bool {
	return m.Control()&CTRLPstuck > 0
}

// Spell is a stringer

func (s *Spell) String() string {
	return s.Name
}

// Ability is a stringer

func (a *Ability) String() string {
	return a.Spell.String()
}
