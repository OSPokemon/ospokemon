package save

import (
	"github.com/ospokemon/ospokemon/engine"
	"time"
)

type Action struct {
	SpellId uint
	Timer   *time.Duration
}

func (a Action) Update(u *engine.Universe, e *engine.Entity, d time.Duration) {
	if a.Timer != nil {
		if *a.Timer < d {
			a.Timer = nil
		} else {
			*a.Timer -= d
		}
	}
}

func (a Action) Snapshot() map[string]interface{} {
	timebuff := 0
	if a.Timer != nil {
		timebuff = int(*a.Timer)
	}

	imagebuff := ""
	if Spells[a.SpellId] != nil {
		imagebuff = Spells[a.SpellId].Image
	}

	return map[string]interface{}{
		"spellid": a.SpellId,
		"image":   imagebuff,
		"timer":   timebuff,
	}
}
