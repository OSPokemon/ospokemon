package persistence

import (
	"ospokemon.com"
	"ztaylor.me/log"
)

func ClassesEntitiesSelect(universe *ospokemon.Universe) (map[uint]*ospokemon.Class, error) {
	rows, err := Connection.Query(
		"SELECT entity, class FROM classes_entities WHERE universe=?",
		universe.Id,
	)
	if err != nil {
		return nil, err
	}

	classes := make(map[uint]*ospokemon.Class)

	for rows.Next() {
		var entityId, classbuff uint
		err := rows.Scan(&entityId, &classbuff)
		if err != nil {
			return nil, err
		}

		class, err := ospokemon.GetClass(classbuff)
		if err != nil {
			return nil, err
		}

		classes[entityId] = class
	}

	log.Add("Classes", classes).Debug("classes_entities select")

	return classes, nil
}
