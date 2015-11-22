package spells

import (
	"github.com/ospokemon/ospokemon/world"
	"log"
	"time"
)

func init() {
	Scripts["Psychic"] = Psychic
}

func Psychic(self world.Entity, t interface{}) {
	target, ok := t.(world.Entity)

	if !ok {
		log.Printf("SpellScript.Psychic (source:%s) invalid target: %v", self.Name(), t)
		return
	}

	effect := &world.Effect{"Psychic Damage", world.EFCThealth, -80, time.Now(), 0}
	target.SetEffects(append(target.Effects(), effect))
}
