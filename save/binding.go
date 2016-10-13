package save

import (
	"github.com/ospokemon/ospokemon/engine"
	"time"
)

type Binding struct {
	Key     string
	SpellId uint
	Timer   *time.Duration
}

func (b Binding) Update(u *engine.Universe, e *engine.Entity, d time.Duration) {
}

func (b Binding) Snapshot() map[string]interface{} {
	timebuff := 0
	if b.Timer != nil {
		timebuff = int(*b.Timer)
	}

	imagebuff := ""
	if Spells[b.SpellId] != nil {
		imagebuff = Spells[b.SpellId].Image
	}

	return map[string]interface{}{
		"key":     b.Key,
		"image":   imagebuff,
		"spellid": b.SpellId,
		"timer":   timebuff,
	}
}
