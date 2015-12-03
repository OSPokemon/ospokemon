package spellscripts

import (
	"github.com/ospokemon/ospokemon/objects/effectscripts"
	"github.com/ospokemon/ospokemon/world"
	"log"
	"time"
)

func init() {
	Scripts["Psychic"] = Psychic
}

func Psychic(self world.Entity, t interface{}, now time.Time) {
	target, ok := t.(world.Entity)

	if !ok {
		log.Printf("SpellScript.Psychic (source:%s) invalid target: %v", self.Name(), t)
		return
	}

	effect := effectscripts.Health.New(80, now, 0)
	target.SetEffects(append(target.Effects(), effect))
}
