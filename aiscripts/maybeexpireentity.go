package aiscripts

import (
	"github.com/ospokemon/ospokemon/engine"
	"time"
)

type Expiry interface {
	Expiration() time.Time
}

func MaybeExpireEntity(m *engine.Map, e engine.Entity, now time.Time) {
	expiry, ok := e.(Expiry)
	if !ok {
		return
	}

	if expiry.Expiration().Before(now) {
		m.RemoveEntity(*e.EntityId())
	}
}
