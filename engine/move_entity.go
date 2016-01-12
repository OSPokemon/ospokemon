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
		for _, e2 := range m.Entities {
			if e == e2 {
				continue
			}
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
