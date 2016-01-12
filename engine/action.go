package engine

import (
	"time"
)

type Action struct {
	Ability      *Ability
	CastStart    *time.Time
	ChannelStart *time.Time
	Target       interface{}
}
