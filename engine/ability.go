package engine

import (
	"time"
)

type Ability struct {
	Spell       *Spell
	Keys        string
	CastTime    time.Duration
	MoveCast    bool
	ChannelTime time.Duration
	MoveChannel bool
	Cooldown    time.Duration
	LastCast    *time.Time
	ManaCost    int
	ItemCost    map[int]int
	TargetType  string
	Range       int
	Size        int
}
