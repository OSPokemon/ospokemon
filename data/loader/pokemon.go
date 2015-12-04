package data

import (
	"github.com/ospokemon/api-go"
	"github.com/ospokemon/ospokemon/data"
	"github.com/ospokemon/ospokemon/objects/entities"
	"github.com/ospokemon/ospokemon/world"
)

func FullLoadPokemon(id int) int {
	pokemon := data.PokemonStore.Load(id)
	pokemon.GRAPHICS = data.GraphicsStore.New("pokemon", pokemon.Species())
	pokemon.CONTROLS = data.ControlsStore.BuildForPokemon(id)

	return world.AddEntity(pokemon)
}

func FullNewAiPokemon(species int, profile *entities.AiProfile) int {
	pokemon := &entities.AiPokemonEntity{
		Entity: data.PokemonEntity{
			BasicPokemon: ospokemon.BasicPokemon{
				SPECIES: species,
			},
			PHYSICS: &world.Physics{
				Position: world.Position{
					X: profile.HomePosition.X,
					Y: profile.HomePosition.Y,
				},
				Size:  world.Size{64, 64},
				Solid: true,
			},
			GRAPHICS: data.GraphicsStore.New("pokemon", species),
			EFFECTS:  make([]*world.Effect, 0),
		},
		Profile: profile,
	}

	pokemon.Entity.CONTROLS = &world.Controls{} // TODO

	return world.AddEntity(pokemon)
}

func FullUnloadPokemon(id int) {
	// TODO
}
