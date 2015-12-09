package loader

import (
	log "github.com/Sirupsen/logrus"
	// "github.com/ospokemon/ospokemon/world"
	// "github.com/ospokemon/ospokemon/eventbus"
	// "github.com/ospokemon/ospokemon/objects/entities"
)

func fetchPokemonInPlayerBox(player_id int, box int) []int {
	rows, err := Connection.Query("SELECT pokemon_id FROM players_pokemon WHERE player_id=? AND box=?", player_id, box)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	pokemon_ids := make([]int, 0)

	var pokemon_id int
	for rows.Next() {
		err = rows.Scan(&pokemon_id)

		if err != nil {
			log.Fatal(err)
		}

		pokemon_ids = append(pokemon_ids, pokemon_id)
	}

	return pokemon_ids
}
