// package data

// import (
// 	log "github.com/Sirupsen/logrus"
// 	"github.com/ospokemon/api-go"
// )

// type speciesStore byte

// var SpeciesStore speciesStore
// var Species = make(map[int]ospokemon.Species)

// func (s speciesStore) Load(id int) ospokemon.Species {
// 	if Species[id] != nil {
// 		return Species[id]
// 	}

// 	importSpecies(id)
// 	importSpeciesTypesRows(id)
// 	importSpeciesSpellsRows(id)

// 	return Species[id]
// }

// func importSpecies(id int) {
// 	row := Connection.QueryRow("SELECT * FROM species WHERE id=?", id)
// 	species := &ospokemon.BasicSpecies{}

// 	err := row.Scan(&species.ID, &species.NAME, &species.TAG, &species.DESCRIPTION, &species.HIDDENABILITY, &species.GENDERLESS, &species.GENDERRATIO, &species.CATCHRATE, &species.BREEDABLE, &species.EGGCYCLES, &species.HEIGHT, &species.WEIGHT, &species.EXPERIENCEYIELD, &species.EXPERIENCECURVE, &species.BODYSTYLE, &species.COLOR, &species.TAMENESS)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	Species[id] = species
// }

// func importSpeciesTypesRows(id int) {
// 	rows, err := Connection.Query("SELECT type_id from species_types where species_id=?", id)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	species := Species[id]
// 	if species.Types() == nil {
// 		species.SetTypes(make([]int, 0))
// 	}

// 	for rows.Next() {
// 		var type_id int
// 		err := rows.Scan(&type_id)

// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		species.SetTypes(append(species.Types(), type_id))
// 	}
// }

// func importSpeciesSpellsRows(id int) {
// 	rows, err := Connection.Query("SELECT spell_id, ai_usable from species_spells where species_id=?", id)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	species := Species[id]
// 	if species.MoveLists() == nil {
// 		species.SetMoveLists(make(map[int][]int))
// 		species.MoveLists()[0] = make([]int, 0)
// 	}
// 	if species.MachineMoves() == nil {
// 		species.SetMachineMoves(make([]int, 0))
// 	}

// 	for rows.Next() {
// 		var spell_id, ai_usable int
// 		err := rows.Scan(&spell_id, &ai_usable)

// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		if ai_usable > 0 {
// 			species.SetMachineMoves(append(species.MachineMoves(), spell_id))
// 		} else {
// 			species.MoveLists()[0] = append(species.MoveLists()[0], spell_id)
// 		}
// 	}
// }
