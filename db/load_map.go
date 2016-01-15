package db

import (
	"github.com/ospokemon/ospokemon/engine"
)

func LoadMap(mapId string) (*engine.Map, error) {
	return &engine.Map{
		Name:     mapId,
		Entities: make([]int, 0),
		Clients:  make([]int, 0),
	}, nil
}
