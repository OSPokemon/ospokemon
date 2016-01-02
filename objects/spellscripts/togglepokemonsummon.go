package spellscripts

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/physics"
	"github.com/ospokemon/ospokemon/registry"
	"github.com/ospokemon/ospokemon/world"
	"strconv"
	"time"
)

func init() {
	registry.Scripts["TogglePokemonSummon"] = TogglePokemonSummon
}

func TogglePokemonSummon(self world.Entity, t interface{}, now time.Time) {
	spell := self.Action().Ability.Spell

	if spell.TargetData["EntityId"] == nil {
		pokemonId, _ := strconv.ParseInt(spell.TargetData["PokemonId"].(string), 10, 0)
		pokemon := registry.Pokemon[int(pokemonId)]
		entityId := world.AddEntity(pokemon)
		spell.TargetData["EntityId"] = entityId
		spell.TargetType = world.TRGTnone

		newposition := t.(physics.Point)
		oldrect := pokemon.Physics().Shape.(physics.Rect)
		newrotation := physics.Vector{1, 0}
		newrect := physics.Rect{newposition, newrotation, oldrect.Height, oldrect.Width}
		pokemon.Physics().Shape = newrect

		go self.(world.Eventer).Fire("SummonPokemon", entityId)

		log.WithFields(log.Fields{
			"PokemonId": pokemonId,
			"EntityId":  entityId,
			"Pokemon":   pokemon.Name(),
		}).Debug("Added Pokemon to world")

	} else {
		entityId, _ := spell.TargetData["EntityId"].(int)
		world.RemoveEntity(entityId)
		delete(spell.TargetData, "EntityId")
		spell.TargetType = world.TRGTposition

		go self.(world.Eventer).Fire("DismissPokemon", entityId)

		log.WithFields(log.Fields{
			"EntityId": entityId,
		}).Debug("Removed Pokemon from world")

	}
}
