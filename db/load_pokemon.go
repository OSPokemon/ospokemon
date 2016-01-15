package db

import (
	"github.com/ospokemon/api-go"
	"github.com/ospokemon/ospokemon/engine"
	"github.com/ospokemon/ospokemon/objects"
	"time"
)

func LoadPokemon(pokemonId int) (*objects.Pokemon, error) {
	row := Connection.QueryRow("SELECT id, name, species, xp, trainer, item FROM pokemon WHERE id=?", pokemonId)

	pokemon := &objects.Pokemon{
		BasicPokemon: ospokemon.MakeBasicPokemon("", 0),
		STATS:        make(map[string]*engine.Stat),
		COLLISION:    engine.CLSNfluid,
	}
	err := row.Scan(&pokemon.ID, &pokemon.NAME, &pokemon.SPECIES, &pokemon.EXPERIENCE, &pokemon.ORIGINALTRAINER, &pokemon.ITEM)
	if err != nil {
		return nil, err
	}

	rows, err := Connection.Query("SELECT stat, value, regen, regenbase, max, base FROM pokemon_stats WHERE pokemonid=?", pokemonId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var statName string
		stat := &engine.Stat{}
		err = rows.Scan(&statName, &stat.Value, &stat.Regen, &stat.RegenBase, &stat.Max, &stat.Base)
		if err != nil {
			return nil, err
		}
	}

	rows, err = Connection.Query("SELECT keys, lastcast, spellid FROM pokemon_abilities WHERE pokemonid=?", pokemonId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var t, spellId int
		ability := &engine.Ability{ItemCost: make(map[int]int)}
		err = rows.Scan(&ability.Keys, &t, &spellId)
		lastcast := time.Unix(int64(t), 0)
		ability.LastCast = &lastcast
		if err != nil {
			return nil, err
		}
		ability.Spell = engine.GetSpell(spellId)
	}

	return pokemon, nil
}
