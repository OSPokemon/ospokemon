package spellscripts

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/objects/effectscripts"
	"github.com/ospokemon/ospokemon/registry"
	"github.com/ospokemon/ospokemon/world"
	"time"
)

func init() {
	registry.Scripts["Psychic"] = Psychic
}

func Psychic(self world.Entity, t interface{}, now time.Time) {
	target, ok := t.(world.Mortality)

	if !ok {
		log.WithFields(log.Fields{
			"source": self.Name(),
			"target": target.Name(),
		}).Error("spellscripts.Psychic invalid target")
		return
	}

	log.WithFields(log.Fields{
		"source": self.Name(),
		"target": target,
	}).Debug("spellscripts.Psychic execution")

	effect := effectscripts.Health.New(80, 0)
	target.SetEffects(append(target.Effects(), effect))
}
