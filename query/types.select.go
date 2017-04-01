package query

import (
	"ospokemon.com/game"
	"ospokemon.com/log"
)

func TypesSelect(id uint) (*game.Type, error) {
	row := Connection.QueryRow(
		"SELECT image FROM types WHERE id=?",
		id,
	)

	t := game.MakeType(id)
	err := row.Scan(&t.Image)

	if err != nil {
		return t, err
	}

	rows, err := Connection.Query(
		"SELECT type2 FROM type_advantage WHERE type1=?",
		id,
	)
	if err != nil {
		return t, err
	}

	for rows.Next() {
		var type2buff uint

		if err := rows.Scan(&type2buff); err != nil {
			return t, err
		}

		t.Strong = append(t.Strong, type2buff)
	}
	rows.Close()

	game.Types[id] = t

	log.Add("Type", id).Info("types select")

	return t, nil
}
