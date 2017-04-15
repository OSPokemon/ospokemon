package persistence

import (
	"ospokemon.com"
)

func init() {
	ospokemon.Universes.Select = UniversesSelect
}

func UniversesSelect(id uint) (*ospokemon.Universe, error) {
	row := Connection.QueryRow(
		"SELECT dx, dy, private FROM universes WHERE id=?",
		id,
	)

	universe := ospokemon.MakeUniverse(id)
	err := row.Scan(&universe.Space.Dimension.DX, &universe.Space.Dimension.DY, &universe.Private)
	if err != nil {
		return nil, err
	}

	entities, err := EntitiesUniversesSelect(universe)
	if err != nil {
		return nil, err
	}
	for _, entity := range entities {
		universe.Add(entity)
	}

	ospokemon.Multiverse[id] = universe

	return universe, nil
}
