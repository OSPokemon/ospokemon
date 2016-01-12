package engine

import (
	"time"
)

var SpellScripts = make(map[string]SpellScript)

type SpellScript func(*Map, Entity, interface{}, time.Time)
