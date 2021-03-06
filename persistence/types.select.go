package persistence

import (
	"github.com/ospokemon/ospokemon"
)

func TypesSelect(id uint) (*ospokemon.Type, error) {
	row := Connection.QueryRow(
		"SELECT image FROM types WHERE id=?",
		id,
	)

	t := ospokemon.MakeType(id)
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

	ospokemon.Types[id] = t

	ospokemon.LOG().Add("Type", id).Info("types select")

	return t, nil
}
