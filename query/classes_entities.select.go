package query

import (
	"ospokemon.com/game"
	"ospokemon.com/log"
)

func ClassesEntitiesSelect(universe *game.Universe) (map[uint]*game.Class, error) {
	rows, err := Connection.Query(
		"SELECT entity, class FROM classes_entities WHERE universe=?",
		universe.Id,
	)
	if err != nil {
		return nil, err
	}

	classes := make(map[uint]*game.Class)

	for rows.Next() {
		var entityId, classbuff uint
		err := rows.Scan(&entityId, &classbuff)
		if err != nil {
			return nil, err
		}

		class, err := GetClass(classbuff)
		if err != nil {
			return nil, err
		}

		classes[entityId] = class
	}

	log.Add("Classes", classes).Debug("classes_entities select")

	return classes, nil
}
