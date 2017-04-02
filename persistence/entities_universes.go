package persistence

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/space"
)

func EntitiesUniversesSelect(universe *ospokemon.Universe) (map[uint]*ospokemon.Entity, error) {
	rows, err := Connection.Query(
		"SELECT id, universe, x, y, dx, dy FROM entities_universes WHERE universe=?",
		universe.Id,
	)

	if err != nil {
		return nil, err
	}

	entities := make(map[uint]*ospokemon.Entity)

	for rows.Next() {
		entity := ospokemon.MakeEntity()
		r := entity.Shape.(*space.Rect)

		err = rows.Scan(&entity.Id, &entity.UniverseId, &r.Anchor.X, &r.Anchor.Y, &r.Dimension.DX, &r.Dimension.DY)
		if err != nil {
			return entities, err
		}
		entities[entity.Id] = entity
	}
	rows.Close()

	event.Fire(event.EntitiesUniversesSelect, entities, universe)
	for _, entity := range entities {
		entity.Id = 0 // delete temp id
	}

	return entities, nil
}
