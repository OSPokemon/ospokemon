package ospokemon

import (
	"time"

	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/space"
	"taylz.io/types"
)

const PARTentity = "entity"

type Entity struct {
	Id         uint
	UniverseId uint
	space.Shape
	Parts
}

type Entities map[uint]*Entity

func MakeEntity() *Entity {
	entity := &Entity{
		Shape: &space.Rect{},
		Parts: make(Parts),
	}

	return entity
}

func (e *Entity) SetClass(c *Class) {
	r := e.Shape.(*space.Rect)
	r.Dimension.DX = c.Dimension.DX
	r.Dimension.DY = c.Dimension.DY

	e.AddPart(BuildImaging(c.Animations))
}

func (e *Entity) Part() string {
	return PARTentity
}

func (parts Parts) GetEntity() *Entity {
	entity, _ := parts[PARTentity].(*Entity)
	return entity
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

func (entity *Entity) Json() types.Dict {
	json := types.Dict{
		"id":    entity.Id,
		"shape": entity.Shape.Snapshot(),
	}

	if imaging := entity.GetImaging(); imaging != nil {
		json["image"] = imaging.Image
	}
	if player := entity.GetPlayer(); player != nil {
		json["player"] = player.Json()
	}
	if chatmessage := entity.GetChatMessage(); chatmessage != nil {
		json["chat"] = chatmessage.Message
	}

	return json
}
