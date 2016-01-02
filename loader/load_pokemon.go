package loader

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ospokemon/api-go"
	"github.com/ospokemon/ospokemon/objects/entities"
	"github.com/ospokemon/ospokemon/physics"
	"github.com/ospokemon/ospokemon/registry"
	"github.com/ospokemon/ospokemon/world"
)

func LoadPokemon(pokemonId int) *entities.PokemonEntity {
	if registry.Pokemon[pokemonId] == nil {
		pokemon := &entities.PokemonEntity{
			BasicPokemon: ospokemon.BasicPokemon{
				STATS: make(map[string]ospokemon.Stat),
			},
			STATHANDLES: make(map[string]world.Stat),
		}
		rect := physics.Rect{physics.Point{}, physics.Vector{1, 0}, 64, 64}

		row := Connection.QueryRow("SELECT id, name, x, y, species, level, experience, ability, friendship, gender, nature, height, weight, originaltrainer, shiny, item FROM pokemon WHERE id=?", pokemonId)
		err := row.Scan(&pokemon.ID, &pokemon.NAME, &rect.Anchor.X, &rect.Anchor.Y, &pokemon.SPECIES, &pokemon.LEVEL, &pokemon.EXPERIENCE, &pokemon.ABILITY, &pokemon.FRIENDSHIP, &pokemon.GENDER, &pokemon.NATURE, &pokemon.HEIGHT, &pokemon.WEIGHT, &pokemon.ORIGINALTRAINER, &pokemon.SHINY, &pokemon.ITEM)
		if err != nil {
			log.Fatal(err)
		}

		pokemon.PHYSICS = &world.Physics{rect, true}
		loadPokemonGraphics(pokemon)
		loadPokemonAbilities(pokemon)
		loadPokemonStats(pokemon)

		registry.Pokemon[pokemonId] = pokemon
	}

	return registry.Pokemon[pokemonId]
}

func UnloadPokemon(entity world.Entity) {
	pokemon := entity.(*entities.PokemonEntity)

	delete(registry.Pokemon, pokemon.Id())
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

func loadPokemonAbilities(pokemon *entities.PokemonEntity) {
	rows, err := Connection.Query("SELECT spell_id, keybinding FROM pokemon_spells WHERE pokemon_id=?", pokemon.Id())
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	pokemon.ABILITIES = make(map[string]*world.Ability)

	var spell_id int
	var keybinding string
	for rows.Next() {
		err = rows.Scan(&spell_id, &keybinding)
		if err != nil {
			log.Fatal(err)
		}

		pokemon.Abilities()[keybinding] = &world.Ability{
			Spell: Spells[spell_id],
		}
	}
}

func loadPokemonStats(pokemon *entities.PokemonEntity) {
	rows, err := Connection.Query("SELECT stat, ev, iv, value FROM pokemon_stats WHERE pokemon_id=?", pokemon.Id())
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var stat_name string
		stat := &ospokemon.BasicStat{}
		err = rows.Scan(&stat_name, &stat.EV, &stat.IV, &stat.VALUE)

		pokemon.BasicPokemon.STATS[stat_name] = stat
	}
}
