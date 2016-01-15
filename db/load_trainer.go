package db

import (
	"github.com/ospokemon/api-go"
	"github.com/ospokemon/ospokemon/engine"
	"github.com/ospokemon/ospokemon/objects"
	"github.com/ospokemon/ospokemon/physics"
)

func LoadTrainer(trainerId int) (*objects.Trainer, error) {
	trainer := &objects.Trainer{
		BasicTrainer: ospokemon.MakeBasicTrainer("", 0),
		STATS:        make(map[string]*engine.Stat),
		COLLISION:    engine.CLSNfluid,
		GRAPHICS:     make(map[engine.AnimationType]string),
	}
	rect := physics.Rect{physics.Point{}, physics.Vector{1, 0}, 64, 64}

	row := Connection.QueryRow("SELECT id, username, name, class, money, x, y, map FROM trainers WHERE id=?", trainerId)
	err := row.Scan(&trainer.BasicTrainer.ID, &trainer.ACCOUNTNAME, &trainer.NAME, &trainer.CLASS, &trainer.MONEY, &rect.Anchor.X, &rect.Anchor.Y, &trainer.MAP)
	if err != nil {
		return nil, err
	}

	trainer.SHAPE = rect

	// loadTrainerGraphics
	class := objects.GetClass(trainer.Class())
	for animationType, image := range class.Graphics {
		trainer.GRAPHICS[animationType] = image
	}
	*trainer.Graphic() = trainer.Graphics()[engine.ANIMwalk_down]

	// loadTrainerGraphics(trainer)
	// loadTrainerStats(trainer)
	// loadTrainerPokemon(trainer)
	// loadTrainerAbilities(trainer)

	return trainer, nil
}

func loadTrainerGraphics(trainer *objects.Trainer) {
	// LoadAnimations("trainer", trainer.Class())
	// trainer.GRAPHICS = Animations["trainer"][trainer.Class()]
	trainer.GRAPHIC = trainer.GRAPHICS[engine.ANIMwalk_down]
}

func loadTrainerStats(trainer *objects.Trainer) {
	rows, _ := Connection.Query("SELECT stat, value, maxvalue, basemaxvalue FROM trainers_stats WHERE trainer_id=?", trainer.Id())
	defer rows.Close()

	trainer.STATS = make(map[string]*engine.Stat)

	var name string
	var stat *engine.Stat
	for rows.Next() {
		stat = &engine.Stat{}
		rows.Scan(&name, &stat.Value, &stat.Regen, &stat.Max, &stat.Base)

		trainer.STATS[name] = stat
	}
}

func unloadTrainerStats(trainer *objects.Trainer) {
	for name, stat := range trainer.Stats() {
		Connection.Exec("UPDATE trainers_stats SET value=?, regen=?, regenbase=?, max=?, base=? WHERE trainer_id=? AND stat=?", stat.Value, stat.Regen, stat.RegenBase, stat.Max, stat.Base, trainer.Id(), name)
	}
}

// func loadTrainerAbilities(trainer *objects.Trainer) {
// 	trainer.ABILITIES = make([]*engine.Ability, 0)

// 	for keybinding, spell := range MakeSpellsForTrainer(trainer.Id()) {
// 		trainer.ABILITIES = append(trainer.ABILITIES, &engine.Ability{
// 			Keys:  keybinding,
// 			Spell: spell,
// 		})
// 	}
// }
