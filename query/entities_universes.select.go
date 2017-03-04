package query

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/space"
)

func EntitiesUniversesSelect(universe *game.Universe) ([]*game.Entity, error) {
	rows, err := Connection.Query(
		"SELECT id, universe, x, y, dx, dy FROM entities_universes WHERE universe=?",
		universe.Id,
	)

	if err != nil {
		return nil, err
	}

	entities := make([]*game.Entity, 0)

	for rows.Next() {
		entity := game.MakeEntity()
		r := entity.Shape.(*space.Rect)

		err = rows.Scan(&entity.Id, &entity.UniverseId, &r.Anchor.X, &r.Anchor.Y, &r.Dimension.DX, &r.Dimension.DY)
		if err != nil {
			return entities, err
		}

		event.Fire(event.EntitiesUniversesSelect, entity, universe)
		entity.Id = 0 // delete temp id
		entities = append(entities, entity)
	}
	rows.Close()

	return entities, nil
}
