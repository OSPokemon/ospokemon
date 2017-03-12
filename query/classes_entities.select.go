package query

import (
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
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
			log.Add("Universe", universe.Id).Add("Entity", entity.Id).Add("Class", classbuff).Debug("classes_entities select")
		}
	}

	return class, err
}
