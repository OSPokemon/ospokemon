package query

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
)

func UniversesSelect(id uint) (*game.Universe, error) {
	row := Connection.QueryRow(
		"SELECT dx, dy, private FROM universes WHERE id=?",
		id,
	)

	universe := game.MakeUniverse(id)
	err := row.Scan(&universe.Space.Dimension.DX, &universe.Space.Dimension.DY, &universe.Private)

	if err == nil {
		log.Add("Universe", "2").Info("universes select")

		event.Fire(event.UniversesSelect, universe)
	}

	game.Multiverse[id] = universe

	return universe, nil
}
