package save

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/engine"
	"time"
)

type Action struct {
	SpellId uint
	Timer   *time.Duration
}

func (a *Action) Update(u *engine.Universe, e *engine.Entity, d time.Duration) {
	if a.Timer == nil {
		return
	}

	if spell := Spells[a.SpellId]; spell != nil {
		if *a.Timer > spell.Cooldown && (*a.Timer-d) <= spell.Cooldown {
			if script := engine.Scripts[spell.ScriptId]; script != nil {
				script(u, e, spell.Data)
			} else {
				logrus.WithFields(logrus.Fields{
					"SpellId": a.SpellId,
				}).Warn("save.Action: Script lookup failed")
			}
		}
	} else {
		logrus.WithFields(logrus.Fields{
			"SpellId": a.SpellId,
		}).Warn("save.Action: Spell lookup failed")
	}

	if *a.Timer < d {
		a.Timer = nil
	} else {
		*a.Timer -= d
	}
}

func (a *Action) Snapshot() map[string]interface{} {
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
