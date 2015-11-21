package data

import (
	"github.com/ospokemon/api-go"
	"log"
)

type classStore byte

var ClassStore classStore
var Classes = make(map[int]ospokemon.Class)

func (c *classStore) Load(id int) ospokemon.Class {
	if Classes[id] != nil {
		return Classes[id]
	}

	row := Connection.QueryRow("SELECT id, name, description FROM classes WHERE id=?", id)
	class := &ospokemon.BasicClass{}

	err := row.Scan(&class.ID, &class.NAME, &class.DESCRIPTION)

	if err != nil {
		log.Fatal(err)
	}

	Classes[id] = class
	return class
}
