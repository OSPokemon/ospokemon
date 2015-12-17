package spellscripts

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/objects/effectscripts"
	"github.com/ospokemon/ospokemon/objects/entities"
	"github.com/ospokemon/ospokemon/physics"
	"github.com/ospokemon/ospokemon/registry"
	"github.com/ospokemon/ospokemon/world"
	"time"
)

func init() {
	registry.Scripts["Gust"] = Gust
}

func Gust(self world.Entity, t interface{}, now time.Time) {
	startpoint := self.Physics().Shape.(*physics.Rect).Anchor
	endpoint := *t.(*physics.Point)

	line := physics.Line{startpoint, endpoint}
	vector := line.Vector()

	rect := physics.Rect{
		Anchor:   startpoint,
		Rotation: vector,
		Height:   physics.DistancePointPoint(startpoint, endpoint),
		Width:    16,
	}

	log.WithFields(log.Fields{
		"Bounds":   rect.MakeLines(),
		"Rotation": vector,
	}).Debug("Gust Wind created")

	expiration := now.Add(2 * time.Second)

	apply := GustEffectApplicator(self, vector.MakeUnit().Multiply(5))

	entity := entities.NewEffectApplicator("Gust Wind", rect, apply, "graphic.gif", expiration)

	world.AddEntity(entity)

}

func GustEffectApplicator(self world.Entity, vector physics.Vector) func(world.Entity) {
	return func(e world.Entity) {
		if e.EntityId() == self.EntityId() {
			return
		}

		mortality, ok := e.(world.Mortality)
		if !ok {
			return
		}

		effect := effectscripts.NewMoveEffect("Gust Wind", vector)

		mortality.SetEffects(append(mortality.Effects(), effect))
	}
}
