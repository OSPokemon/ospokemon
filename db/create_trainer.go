package db

import (
	"github.com/ospokemon/api-go"
	"github.com/ospokemon/ospokemon/engine"
	"github.com/ospokemon/ospokemon/objects"
	"github.com/ospokemon/ospokemon/physics"
	"strings"
)

func CreateTrainer(username, name string, classId int) (*objects.Trainer, error) {
	trainer := &objects.Trainer{
		BasicTrainer: ospokemon.MakeBasicTrainer(name, classId),
		STATS:        make(map[string]*engine.Stat),
		COLLISION:    engine.CLSNfluid,
		SHAPE:        physics.Rect{physics.Point{}, physics.Vector{1, 0}, 64, 64},
		GRAPHICS:     make(map[engine.AnimationType]string),
	}

	res, err := Connection.Exec("INSERT INTO trainers (username, name, class, money, map, x, y) VALUES (?, ?, ?, ?, ?, ?, ?)", username, trainer.Name(), trainer.Class(), trainer.Money(), "global", 0, 0)
	if err != nil {
		return nil, err
	}

	trainerId, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	trainer.SetId(int(trainerId))

	// loadTrainerGraphics
	class := objects.GetClass(trainer.Class())
	for animationType, image := range class.Graphics {
		trainer.GRAPHICS[animationType] = image
	}

	// loadTrainerStats
	for statName, val := range class.Stats {
		if strings.HasSuffix(statName, "-regen") {
			statName = strings.TrimSuffix(statName, "-regen")

			if stat := trainer.STATS[statName]; stat != nil {
				stat.RegenBase = val
			} else {
				trainer.STATS[statName] = &engine.Stat{
					RegenBase: val,
				}
			}
		} else if stat := trainer.STATS[statName]; stat != nil {
			stat.Base = val
		} else {
			trainer.STATS[statName] = &engine.Stat{
				Base: val,
			}
		}
	}

	// loadTrainerGraphics(trainer)
	// loadTrainerPokemon(trainer)
	// loadTrainerAbilities(trainer)

	return trainer, nil
}
