// package loader

// import (
// 	"github.com/ospokemon/api-go"
// 	"github.com/ospokemon/ospokemon/data"
// 	"github.com/ospokemon/ospokemon/objects/entities"
// 	"github.com/ospokemon/ospokemon/world"
// )

// func FullLoadPokemon(id int) int {
// 	pokemon := data.PokemonStore.Load(id)
// 	pokemon.GRAPHICS = data.GraphicsStore.New("pokemon", pokemon.Species())
// 	pokemon.ABILITIES = data.AbilitiesStore.GetForPokemon(id)

// 	return world.AddEntity(pokemon)
// }

// func FullNewAiPokemon(speciesId int, profile *entities.AiProfile) int {
// 	pokemon := &entities.AiPokemonEntity{
// 		Entity: data.PokemonEntity{
// 			BasicPokemon: ospokemon.BasicPokemon{
// 				SPECIES: speciesId,
// 				STATS:   make(map[string]ospokemon.Stat),
// 			},
// 			PHYSICS: &world.Physics{
// 				Point: world.Point{
// 					X: profile.HomePoint.X,
// 					Y: profile.HomePoint.Y,
// 				},
// 				Size:  world.Size{64, 64},
// 				Solid: true,
// 			},
// 			GRAPHICS:    data.GraphicsStore.New("pokemon", speciesId),
// 			STATHANDLES: make(map[string]world.Stat),
// 			ABILITIES:   make(map[string]*world.Ability),
// 			EFFECTS:     make([]*world.Effect, 0),
// 		},
// 		Profile: profile,
// 	}

// 	species := data.SpeciesStore.Load(speciesId)

// 	pokemon.Entity.BasicPokemon.Stats()["health"] = &ospokemon.BasicStat{
// 		IV:    species.Stats()["health"],
// 		EV:    species.Stats()["health"],
// 		VALUE: species.Stats()["health"],
// 	}
// 	pokemon.Entity.BasicPokemon.Stats()["speed"] = &ospokemon.BasicStat{
// 		IV:    species.Stats()["speed"],
// 		EV:    species.Stats()["speed"],
// 		VALUE: species.Stats()["speed"],
// 	}

// 	pokemon.Entity.ABILITIES = make(map[string]*world.Ability) // TODO

// 	return world.AddEntity(pokemon)
// }

// func FullUnloadPokemon(id int) {
// 	// TODO
// }
