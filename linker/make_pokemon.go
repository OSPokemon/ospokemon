package linker

import (
	"github.com/ospokemon/api-go"
	"github.com/ospokemon/ospokemon/engine"
	"github.com/ospokemon/ospokemon/objects"
	"github.com/ospokemon/ospokemon/physics"
	"strings"
)

func MakePokemon(name string, speciesId int) (*objects.Pokemon, error) {
	pokemon := &objects.Pokemon{
		BasicPokemon: ospokemon.MakeBasicPokemon(name, speciesId),
		COLLISION:    engine.CLSNfluid,
		STATS:        make(map[string]*engine.Stat),
		SHAPE:        physics.Rect{physics.Point{}, physics.Vector{1, 0}, 64, 64},
		GRAPHICS:     make(map[engine.AnimationType]string),
	}

	species := objects.GetSpecies(speciesId)

	if pokemon.Name() == "" {
		pokemon.SetName(species.Name())
	}

	for animationType, image := range species.GRAPHICS {
		pokemon.GRAPHICS[animationType] = image
	}

	for statName, val := range species.Stats() {
		if strings.HasSuffix(statName, "-regen") {
			statName = strings.TrimSuffix(statName, "-regen")

			if stat := pokemon.STATS[statName]; stat != nil {
				stat.RegenBase = val
			} else {
				pokemon.STATS[statName] = &engine.Stat{
					RegenBase: val,
				}
			}
		} else if stat := pokemon.STATS[statName]; stat != nil {
			stat.Base = val
		} else {
			pokemon.STATS[statName] = &engine.Stat{
				Base: val,
			}
		}
	}

	return pokemon, nil
}
