package game

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/part"
	"github.com/ospokemon/ospokemon/space"
	"time"
)

type Entity struct {
	Id         uint
	UniverseId uint
	space.Shape
	part.Parts
}

type Entities map[uint]*Entity

func MakeEntity() *Entity {
	entity := &Entity{
		Shape: &space.Rect{},
		Parts: make(part.Parts),
	}

	return entity
}

func (e *Entity) Part() string {
	return part.Entity
}

func (e *Entity) Update(u *Universe, d time.Duration) {
	for _, part := range e.Parts {
		if updater, ok := part.(Updater); ok {
			updater.Update(u, e, d)
		}
	}
}

func (e *Entity) Move(vector space.Vector, universe *Universe) {
	e.Shape = e.Shape.Move(vector)
	event.Fire(event.Movement, e, vector)

	if vector.Length() < 1 {
		return
	}

	for entityId, entity := range universe.Entities {
		if entity == nil {
			continue
		}

		if e.Id == entityId {
			continue
		}

		if space.DistanceShapeShape(e.Shape, entity.Shape) < 1 {
			event.Fire(event.Collision, e, entity, universe, vector)
		}
	}
}
