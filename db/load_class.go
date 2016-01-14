package db

import (
	"github.com/ospokemon/api-go"
	"github.com/ospokemon/ospokemon/engine"
	"github.com/ospokemon/ospokemon/objects"
)

func LoadClass(classId int) (*objects.Class, error) {
	row := Connection.QueryRow("SELECT id, name, description FROM classes WHERE id=?", classId)

	class := &objects.Class{
		BasicClass: ospokemon.BasicClass{},
		Graphics:   make(map[engine.AnimationType]string),
		Stats:      make(map[string]float64),
	}

	err := row.Scan(&class.ID, &class.NAME, &class.DESCRIPTION)
	if err != nil {
		return nil, err
	}

	rows, err := Connection.Query("SELECT stat, value FROM class_stats WHERE classid=?", classId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var stat string
		var value float64
		rows.Scan(&stat, &value)
		class.Stats[stat] = value
	}

	rows, err = Connection.Query("SELECT anim, image FROM class_animations WHERE classid=?", classId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var anim, image string
		rows.Scan(&anim, &image)
		class.Graphics[engine.AnimationType(anim)] = image
	}

	return class, nil
}
