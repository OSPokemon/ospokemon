package db

import (
	"github.com/ospokemon/api-go"
	"github.com/ospokemon/ospokemon/engine"
	"github.com/ospokemon/ospokemon/objects"
	"github.com/ospokemon/ospokemon/physics"
)

func CreateTrainer(username, name string, class int) (*objects.Trainer, error) {
	trainer := &objects.Trainer{
		BasicTrainer: ospokemon.MakeBasicTrainer(name, class),
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

	// trainer graphics
	rows, err := Connection.Query("SELECT anim, image FROM trainer_animations WHERE trainerclass=?", trainer.CLASS)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var anim, image string
		rows.Scan(&anim, &image)
		trainer.GRAPHICS[engine.AnimationType(anim)] = image
	}

	// loadTrainerGraphics(trainer)
	// loadTrainerStats(trainer)
	// loadTrainerPokemon(trainer)
	// loadTrainerAbilities(trainer)

	return trainer, nil
}
