package persistence

import (
	"github.com/ospokemon/ospokemon"
	"github.com/pkg/errors"
)

func ClassesEntitiesSelect(universe *ospokemon.Universe) (map[uint]*ospokemon.Class, error) {
	rows, err := Connection.Query(
		"SELECT entity, class FROM classes_entities WHERE universe=?",
		universe.Id,
	)
	if err != nil {
		return nil, errors.Wrap(err, "classes_entities.select")
	}

	classesData := make(map[uint]uint)

	for rows.Next() {
		var entityId, classbuff uint
		if err := rows.Scan(&entityId, &classbuff); err != nil {
			return nil, errors.Wrap(err, "classes_entities.scan")
		}
		classesData[entityId] = classbuff
	}
	rows.Close()

	classes := make(map[uint]*ospokemon.Class)

	for k, v := range classesData {
		if class, err := ospokemon.GetClass(v); err != nil {
			return nil, errors.Wrap(err, "classes_entities.getclass")
		} else {
			classes[k] = class
		}
	}

	ospokemon.LOG().Add("Classes", classes).Debug("classes_entities select")
	return classes, nil
}
