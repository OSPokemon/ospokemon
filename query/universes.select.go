package query

import (
	"ospokemon.com/event"
	"ospokemon.com/game"
	"ospokemon.com/log"
)

func UniversesSelect(id uint) (*game.Universe, error) {
	row := Connection.QueryRow(
		"SELECT dx, dy, private FROM universes WHERE id=?",
		id,
	)

	universe := game.MakeUniverse(id)
	err := row.Scan(&universe.Space.Dimension.DX, &universe.Space.Dimension.DY, &universe.Private)

	if err == nil {
		log.Add("Universe", id).Info("universes select")

		event.Fire(event.UniversesSelect, universe)
	}

	game.Multiverse[id] = universe

	return universe, nil
}
