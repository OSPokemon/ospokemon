package data

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ospokemon/api-go"
	"github.com/ospokemon/ospokemon/world"
)

type pokemonStore byte

var PokemonStore pokemonStore
var Pokemon = make(map[int]*PokemonEntity)

func (p *pokemonStore) FetchIdsInPlayerBox(player_id int, box int) []int {
	rows, err := Connection.Query("SELECT pokemon_id FROM players_pokemon WHERE player_id=? AND box=?", player_id, box)
	defer rows.Close()

	if err != nil {
		log.Fatal(err)
	}

	pokemon_ids := make([]int, 0)
	var pokemon_id int

	for rows.Next() {
		err = rows.Scan(&pokemon_id)

		if err != nil {
			log.Fatal(err)
		}

		pokemon_ids = append(pokemon_ids, pokemon_id)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return pokemon_ids
}

func (p *pokemonStore) Load(id int) *PokemonEntity {
	if Pokemon[id] != nil {
		return Pokemon[id]
	}

	row := Connection.QueryRow("SELECT id, name, x, y, species, level, experience, ability, friendship, gender, nature, height, weight, originaltrainer, shiny, item FROM pokemon WHERE id=?", id)
	pokemon := &PokemonEntity{
		PHYSICS: &world.Physics{
			Position: world.Position{},
			Size:     world.Size{64, 64},
			Solid:    true,
		},
		STATHANDLES: make(map[string]world.Stat),
		ABILITIES:   make(map[string]*world.Ability),
	}

	err := row.Scan(&pokemon.ID, &pokemon.NAME, &pokemon.PHYSICS.Position.X, &pokemon.PHYSICS.Position.Y, &pokemon.SPECIES, &pokemon.LEVEL, &pokemon.EXPERIENCE, &pokemon.ABILITY, &pokemon.FRIENDSHIP, &pokemon.GENDER, &pokemon.NATURE, &pokemon.HEIGHT, &pokemon.WEIGHT, &pokemon.ORIGINALTRAINER, &pokemon.SHINY, &pokemon.ITEM)

	if err != nil {
		log.Fatal(err)
	}

	rows, err := Connection.Query("SELECT stat, ev, iv, value FROM pokemon_stats WHERE pokemon_id=?", id)

	if err != nil {
		log.Fatal(err)
	}

	pokemon.STATS = make(map[string]ospokemon.Stat)

	for rows.Next() {
		stat_name := ""
		stat := &ospokemon.BasicStat{}
		err = rows.Scan(&stat_name, &stat.EV, &stat.IV, &stat.VALUE)

		pokemon.BasicPokemon.STATS[stat_name] = stat
	}

	Pokemon[id] = pokemon
	return pokemon
}
