package query

import (
	"ospokemon.com/game"
	"ospokemon.com/log"
)

func ClassesSelect(id uint) (*game.Class, error) {
	row := Connection.QueryRow(
		"SELECT dx, dy FROM classes WHERE id=?",
		id,
	)

	class := game.MakeClass(id)
	err := row.Scan(&class.Dimension.DX, &class.Dimension.DY)

	if err != nil {
		return class, err
	}

	rows, err := Connection.Query(
		"SELECT key, value FROM animations_classes WHERE class=?",
		class.Id,
	)

	if err != nil {
		return class, err
	}

	for rows.Next() {
		var keybuff, valuebuff string
		err = rows.Scan(&keybuff, &valuebuff)
		class.Animations[keybuff] = valuebuff
	}
	rows.Close()

	game.Classes[id] = class

	log.Add("Class", class.Id).Info("classes select")

	return class, nil
}
