package aiscripts

import (
	"github.com/ospokemon/ospokemon/world"
	"time"
)

type Expiry interface {
	Expiration() time.Time
}

func MaybeExpireEntity(e world.Entity, now time.Time) {
	expiry, ok := e.(Expiry)
	if !ok {
		return
	}

	if expiry.Expiration().Before(now) {
		world.RemoveEntity(e.EntityId())
	}
}
