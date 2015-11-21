package data

import (
// "database/sql"
// "log"
// "src.ospokemon.org/ospokemon"
)

// var SpeciesStore = make(speciesStore)

// type speciesStore map[int]ospokemon.Species

// func (s speciesStore) Load(id int) ospokemon.Species {
// 	if s[id] != nil {
// 		return s[id]
// 	}

// 	row := Connection.QueryRow("SELECT * FROM species WHERE id=?", id)

// 	importSpeciesRow(row)

// 	rows, _ := Connection.Query("SELECT * from species_types where species_id=?", id)

// 	importSpeciesTypesRows(rows)

// 	return SpeciesStore[id]
// }

// func importSpeciesRow(row *sql.Row) {
// 	species := &ospokemon.BasicSpecies{}

// 	var id, hidden_ability, genderless, catch_rate, breedable, egg_cycles, xp_yield, xp_curve, tameness int
// 	var name, tag, description, body_style, color string
// 	var gender_ratio, height, weight float64

// 	err := row.Scan(&id, &name, &tag, &description, &hidden_ability, &genderless, &gender_ratio, &catch_rate,
// 		&breedable, &egg_cycles, &height, &weight, &xp_yield, &xp_curve, &body_style, &color, &tameness)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	species.SetId(id)
// 	species.SetName(name)
// 	species.SetTag(tag)
// 	species.SetDescription(description)
// 	// TODO abilities
// 	species.SetHiddenAbility(hidden_ability)

// 	if genderless > 0 {
// 		species.SetGenderless(true)
// 	} else {
// 		species.SetGenderless(false)
// 	}

// 	species.SetGenderRatio(gender_ratio)
// 	species.SetCatchRate(catch_rate)

// 	if breedable > 0 {
// 		species.SetBreedable(true)
// 	} else {
// 		species.SetGenderless(false)
// 	}

// 	species.SetEggCycles(egg_cycles)
// 	species.SetHeight(height)
// 	species.SetWeight(weight)
// 	species.SetExperienceYield(xp_yield)
// 	species.SetExperienceCurve(xp_curve)
// 	species.SetBodyStyle(body_style)
// 	species.SetColor(color)
// 	species.SetTameness(tameness)

// 	SpeciesStore[species.Id()] = species
// }

// func importSpeciesTypesRows(rows *sql.Rows) {
// 	for rows.Next() {
// 		var species_id, type_id int
// 		err := rows.Scan(&species_id, &type_id)

// 		if err != nil {
// 			continue
// 		}

// 		species := SpeciesStore[species_id]
// 		species_types := species.Types()

// 		if species_types == nil {
// 			species_types = make([]int, 1)
// 		}

// 		species_types = append(species_types, type_id)
// 		species.SetTypes(species_types)
// 	}
// }
