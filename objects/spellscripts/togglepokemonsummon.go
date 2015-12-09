package spellscripts

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/registry"
	"github.com/ospokemon/ospokemon/world"
	"time"
)

func init() {
	registry.Scripts["TogglePokemonSummon"] = TogglePokemonSummon
}

func TogglePokemonSummon(self world.Entity, t interface{}, now time.Time) {
	log.WithFields(log.Fields{
		"Target": t,
		"Entity": self.Name(),
	}).Warn("TogglePokemonSummon")
}
