package loader

import (
	"github.com/ospokemon/ospokemon/data"
	"github.com/ospokemon/ospokemon/world"
)

func FullLoadPlayer(name string) []int {
	entities := make([]int, 0)

	player := data.PlayerStore.Load(name)
	player.GRAPHICS = data.GraphicsStore.New("trainer", player.Class())
	player.ABILITIES = make(map[string]*world.Ability)

	entities = append(entities, world.AddEntity(player))

	pokemon_ids := data.PokemonStore.FetchIdsInPlayerBox(player.Id(), 0)

	for _, pokemon_id := range pokemon_ids {
		entities = append(entities, FullLoadPokemon(pokemon_id))
	}

	player.BasicTrainer.POKEMON = pokemon_ids

	return entities
}

func FullUnloadPlayer(name string) {
	// todo propogate changes to disk
}