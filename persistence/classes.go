package persistence

import (
	"ospokemon.com"
	"ztaylor.me/log"
)

func init() {
	ospokemon.Classes.Select = ClassesSelect
}

func ClassesSelect(id uint) (*ospokemon.Class, error) {
	row := Connection.QueryRow(
		"SELECT dx, dy FROM classes WHERE id=?",
		id,
	)

	class := ospokemon.MakeClass(id)
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

	log.Add("Class", class.Id).Info("classes select")

	return class, nil
}
