package query

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/game"
)

func ClassesEntitiesSelect(entity *game.Entity, universe *game.Universe) (*game.Class, error) {
	row := Connection.QueryRow(
		"SELECT class FROM classes_entities WHERE entity=? AND universe=?",
		entity.Id,
		universe.Id,
	)

	var classbuff uint
	var class *game.Class
	err := row.Scan(&classbuff)

	if err == nil {
		class, err = GetClass(classbuff)

		if err == nil {
			logrus.WithFields(logrus.Fields{
				"Universe": universe.Id,
				"Entity":   entity.Id,
				"Class":    classbuff,
			}).Debug("classes_entities select")
		}
	}

	return class, err
}
