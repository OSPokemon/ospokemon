package persistence

import (
	"github.com/pkg/errors"
	"github.com/ospokemon/ospokemon"
	"github.com/ospokemon/ospokemon/space"
)

func EntitiesUniversesSelect(universe *ospokemon.Universe) ([]*ospokemon.Entity, error) {
	rows, err := Connection.Query(
		"SELECT id, universe, x, y, dx, dy FROM entities_universes WHERE universe=?",
		universe.Id,
	)

	if err != nil {
		return nil, errors.Wrap(err, "entities_universes.select")
	}

	entities := make(map[uint]*ospokemon.Entity)

	for rows.Next() {
		entity := ospokemon.MakeEntity()
		r := entity.Shape.(*space.Rect)

		err = rows.Scan(&entity.Id, &entity.UniverseId, &r.Anchor.X, &r.Anchor.Y, &r.Dimension.DX, &r.Dimension.DY)
		if err != nil {
			return nil, errors.Wrap(err, "entities_universes.scan")
		}
		entities[entity.Id] = entity
	}
	rows.Close()

	// classes
	classes, err := ClassesEntitiesSelect(universe)
	if err != nil {
		return nil, errors.Wrap(err, "entities_universes")
	}
	for entityId, class := range classes {
		entity := entities[entityId]

		entity.AddPart(ospokemon.BuildImaging(class.Animations))

		rect := entity.Shape.(*space.Rect)
		rect.Dimension.DX = class.Dimension.DX
		rect.Dimension.DY = class.Dimension.DY
	}

	// dialogs
	dialogs, err := DialogsSelect(universe.Id)
	if err != nil {
		return nil, errors.Wrap(err, "entities_universes")
	}
	for entityId, dialog := range dialogs {
		entities[entityId].AddPart(dialog)
	}

	// itemslots
	itemslots, err := EntitiesItemsSelect(universe)
	if err != nil {
		return nil, errors.Wrap(err, "entities_universes")
	}
	for entityId, itemslot := range itemslots {
		entity := entities[entityId]
		entity.AddPart(itemslot)
		entity.AddPart(itemslot.GetImaging())
		itemslot.Parts = entity.Parts

		rect := entity.Shape.(*space.Rect)
		item := itemslot.Item
		rect.Dimension.DX = item.Dimension.DX
		rect.Dimension.DY = item.Dimension.DY
	}

	// spawners
	spawners, err := EntitiesSpawnersSelect(universe)
	if err != nil {
		return nil, errors.Wrap(err, "entities_universes")
	}
	for entityId, spawner := range spawners {
		spawner.Child = entities[entityId]
		spawner.Child.AddPart(spawner)
		universe.AddSpawner(spawner)
	}

	// terrain
	terrains, err := EntitiesTerrainsSelect(universe)
	if err != nil {
		return nil, errors.Wrap(err, "entities_universes")
	}
	for entityId, terrain := range terrains {
		entity := entities[entityId]
		entity.AddPart(terrain)

		imaging := ospokemon.MakeImaging()
		imaging.Image = terrain.Image
		entity.AddPart(imaging)
	}

	// build array to return
	i := 0
	ret := make([]*ospokemon.Entity, len(entities))
	for _, entity := range entities {
		ret[i] = entity
		i++
	}

	return ret, nil
}
