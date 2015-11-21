package data

import (
	"github.com/ospokemon/ospokemon/world"
	"log"
)

// Load all the resources for a Player, binding appropriately, insert into world
func FullLoadPlayer(name string) []int {
	entities := make([]int, 0)
	log.Printf("entities after allocation: %v", entities)

	player := PlayerStore.Load(name)
	player.GRAPHICS = GraphicsStore.New(player.Class())

	player.CONTROLS = &world.Controls{
		Abilities: make(map[string]*world.Ability),
	}

	entities = append(entities, world.AddEntity(player))

	return entities
}
