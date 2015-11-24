package data

import (
	"github.com/ospokemon/ospokemon/world"
)

// Load all the resources for a Player, binding appropriately, insert into world
func FullLoadPlayer(name string) []int {
	entities := make([]int, 0)

	player := PlayerStore.Load(name)
	player.GRAPHICS = GraphicsStore.New("trainerclass", player.Class())
	player.CONTROLS = &world.Controls{
		Abilities: make(map[string]*world.Ability),
	}

	entities = append(entities, world.AddEntity(player))

	pokemon_ids := PokemonStore.FetchIdsInPlayerBox(player.Id(), 0)

	for _, pokemon_id := range pokemon_ids {
		pokemon := PokemonStore.Load(pokemon_id)
		pokemon.GRAPHICS = GraphicsStore.New("pokemon", pokemon.Species())
		pokemon.CONTROLS = ControlsStore.BuildForPokemon(pokemon_id)

		entities = append(entities, world.AddEntity(pokemon))
	}

	return entities
}

func FullUnloadPlayer(name string) {
	// todo propogate changes to disk
}
