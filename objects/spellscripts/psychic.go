package spellscripts

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/objects/effectscripts"
	"github.com/ospokemon/ospokemon/world"
	"time"
)

func init() {
	Scripts["Psychic"] = Psychic
}

func Psychic(self world.Entity, t interface{}, now time.Time) {
	target, ok := t.(world.Entity)

	if !ok {
		log.WithFields(log.Fields{
			"source": self.Name(),
			"target": t,
		}).Error("spellscripts.Psychic invalid target")
		return
	}

	effect := effectscripts.Health.New(80, now, 0)
	target.SetEffects(append(target.Effects(), effect))
}
