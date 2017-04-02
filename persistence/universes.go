package persistence

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
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

	if err == nil {
		log.Add("Universe", id).Info("universes select")

		event.Fire(event.UniversesSelect, universe)
	}

	ospokemon.Multiverse[id] = universe

	return universe, nil
}
