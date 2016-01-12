package engine

import (
	"time"
)

type EffectScript func(effect *Effect, m *Map, entity Entity, now time.Time)

type Effect struct {
	Name     string
	Data     map[string]interface{}
	Start    *time.Time
	Duration time.Duration
	EffectScript
}
