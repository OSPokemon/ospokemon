package engine

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/physics"
)

func MoveEntity(m *Map, e LivingEntity, v physics.Vector) {
	if *e.Control()&CTRLPstuck < 1 {
		return
	}

	nextShape := e.Shape().Move(v)

	if *e.Collision() != CLSNnone {
		for _, entityId2 := range m.Entities {
			if *e.EntityId() == entityId2 {
				continue
			}
			e2 := Entities[entityId2]
			if *e2.Collision() == CLSNnone {
				continue
			}

			if CheckCollision(nextShape, e.Shape()) {
				return
			}
		}
	}

	log.WithFields(log.Fields{
		"Entity": e.Name(),
		"Vector": v,
		"Shape":  nextShape,
	}).Debug("Move entity")

	e.SetShape(nextShape)
}
