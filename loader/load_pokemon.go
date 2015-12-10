package loader

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ospokemon/api-go"
	"github.com/ospokemon/ospokemon/objects/entities"
	"github.com/ospokemon/ospokemon/registry"
	"github.com/ospokemon/ospokemon/world"
)

func init() {
	registry.Loaders["Pokemon"] = LoadPokemon
}

func LoadPokemon(pokemonId int) {
	if registry.Pokemon[pokemonId] != nil {
		return
	}

	row := Connection.QueryRow("SELECT id, name, x, y, species, level, experience, ability, friendship, gender, nature, height, weight, originaltrainer, shiny, item FROM pokemon WHERE id=?", pokemonId)
	pokemon := &entities.PokemonEntity{
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

	rows, err := Connection.Query("SELECT stat, ev, iv, value FROM pokemon_stats WHERE pokemon_id=?", pokemonId)
	if err != nil {
		log.Fatal(err)
	}

	loadPokemonGraphics(pokemon)
	pokemon.STATS = make(map[string]ospokemon.Stat)

	for rows.Next() {
		stat_name := ""
		stat := &ospokemon.BasicStat{}
		err = rows.Scan(&stat_name, &stat.EV, &stat.IV, &stat.VALUE)

		pokemon.BasicPokemon.STATS[stat_name] = stat
	}

	registry.Pokemon[pokemonId] = pokemon

	log.WithFields(log.Fields{
		"Pokemon": pokemon,
	}).Debug("Pokemon built")
}

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

func loadPokemonGraphics(pokemon *entities.PokemonEntity) {
	LoadAnimations("pokemon", pokemon.Species())
	animations := Animations["pokemon"][pokemon.Species()]
	pokemon.GRAPHICS = &world.Graphics{
		Portrait:   animations[world.ANIMportrait],
		Current:    animations[world.ANIMwalk_down],
		Animations: animations,
	}
}
