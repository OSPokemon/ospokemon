package db

import (
	"github.com/ospokemon/ospokemon/objects"
)

func SavePokemon(pokemon *objects.Pokemon) error {

	if pokemon.Id() == 0 {
		if err := insertPokemon(pokemon); err != nil {
			return err
		}
	} else {
		if err := updatePokemon(pokemon); err != nil {
			return err
		}
	}

	Connection.Exec("DELETE FROM pokemon_stats WHERE pokemonid=?", pokemon.Id())
	for statName, stat := range pokemon.Stats() {
		Connection.Exec("INSERT INTO pokemon_stats (pokemonid, stat, value, regen, regenbase, max, base) VALUES (?, ?, ?, ?, ?, ?, ?)", pokemon.Id(), statName, stat.Value, stat.Regen, stat.RegenBase, stat.Max, stat.Base)
	}

	Connection.Exec("DELETE FROM pokemon_abilities WHERE pokemonid=?", pokemon.Id())
	for _, ability := range *pokemon.Abilities() {
		Connection.Exec("INSERT INTO pokemon_abilities (pokemonid, keys, lastcast, spellid) VALUES (?, ?, ?, ?)", pokemon.Id(), ability.Keys, ability.LastCast.Unix(), ability.Spell.Id)
	}

	return nil
}

func insertPokemon(pokemon *objects.Pokemon) error {

	res, err := Connection.Exec("INSERT INTO pokemon (name, species, xp, trainer, item) VALUES (?, ?, ?, ?, ?)", pokemon.Name(), pokemon.Species(), pokemon.Experience(), pokemon.OriginalTrainer(), pokemon.Item())
	if err == nil {
		var id int64
		id, err = res.LastInsertId()
		pokemon.SetId(int(id))
	}
	return err
}

func updatePokemon(pokemon *objects.Pokemon) error {

	_, err := Connection.Exec("UPDATE pokemon SET name=?, xp=?, item=? WHERE id=?", pokemon.Name(), pokemon.Experience(), pokemon.Item(), pokemon.Id())
	return err
}
