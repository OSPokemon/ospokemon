package query

import (
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
)

func SpeciesSelect(id uint) (*game.Species, error) {
	row := Connection.QueryRow(
		"SELECT name, genderratio, catchfactor, hatchsteps, height, width, xpfunc FROM species WHERE id=?",
		id,
	)

	species := game.MakeSpecies(id)

	var genderratiobuff float64
	err := row.Scan(
		&species.Name,
		&genderratiobuff,
		&species.CatchFactor,
		&species.HatchSteps,
		&species.Height,
		&species.Width,
		&species.XpFunc,
	)
	if err != nil {
		return nil, err
	}

	if genderratiobuff >= 0 {
		species.GenderRatio = &genderratiobuff
	}

	// Species Types
	rows, err := Connection.Query(
		"SELECT type FROM species_types WHERE species=?",
		id,
	)
	if err != nil {
		return species, err
	}

	for rows.Next() {
		var typebuff uint
		rows.Scan(&typebuff)
		species.Types = append(species.Types, typebuff)
	}
	rows.Close()

	// Species Mate groups
	rows, err = Connection.Query(
		"SELECT group FROM species_mate_groups WHERE species=?",
		id,
	)
	if err != nil {
		return species, err
	}

	for rows.Next() {
		var groupbuff uint
		rows.Scan(&groupbuff)
		species.MateGroups = append(species.MateGroups, groupbuff)
	}
	rows.Close()

	// Species Level moves
	rows, err = Connection.Query(
		"SELECT level, spell FROM species_level_moves WHERE species=?",
		id,
	)
	if err != nil {
		return species, err
	}

	for rows.Next() {
		var levelbuff, spellbuff uint
		rows.Scan(&levelbuff, &spellbuff)

		if species.LevelMoves[levelbuff] == nil {
			species.LevelMoves[levelbuff] = []uint{
				spellbuff,
			}
		} else {
			species.LevelMoves[levelbuff] = append(species.LevelMoves[levelbuff], spellbuff)
		}
	}
	rows.Close()

	// Species Hatch moves
	rows, err = Connection.Query(
		"SELECT spell FROM species_hatch_moves WHERE species=?",
		id,
	)
	if err != nil {
		return species, err
	}

	for rows.Next() {
		var movebuff uint
		rows.Scan(&movebuff)
		species.HatchMoves = append(species.HatchMoves, movebuff)
	}
	rows.Close()

	// Species Stats
	rows, err = Connection.Query(
		"SELECT stat, value FROM species_stats WHERE species=?",
		id,
	)
	if err != nil {
		return species, err
	}

	for rows.Next() {
		var statbuff uint
		var valuebuff float64
		rows.Scan(&statbuff, &valuebuff)
		species.Stats[statbuff] = valuebuff
	}
	rows.Close()

	// Species Animations
	rows, err = Connection.Query(
		"SELECT key, value FROM animations_species WHERE species=?",
		id,
	)
	if err != nil {
		return species, err
	}

	for rows.Next() {
		var keybuff, valuebuff string
		rows.Scan(&keybuff, &valuebuff)
		species.Animations[keybuff] = valuebuff
	}
	rows.Close()

	log.Add("Species", id).Info("species select")

	return species, nil
}
