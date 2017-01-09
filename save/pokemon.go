package save

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/part"
)

type Pokemon struct {
	Id      uint
	Species uint
	Xp      uint
	Level   uint
	Gender  string
	Shiny   bool
	part.Parts
}

func MakePokemon(id uint) *Pokemon {
	p := &Pokemon{
		Id:    id,
		Parts: make(part.Parts),
	}

	event.Fire(event.PokemonMake, p)

	return p
}

var Pokemons = make(map[uint]*Pokemon)

func GetPokemon(id uint) (*Pokemon, error) {
	if p, ok := Pokemons[id]; p != nil {
		return p, nil
	} else if ok {
		return nil, nil
	}

	p := MakePokemon(id)
	err := p.Query()
	if err != nil {
		p = nil
	}

	Pokemons[id] = p
	return p, err
}

func (p *Pokemon) Query() error {
	row := Connection.QueryRow(
		"SELECT species, xp, level, gender, shiny FROM pokemon WHERE id=?",
		p.Id,
	)

	if err := row.Scan(&p.Species, &p.Xp, &p.Level, &p.Gender, &p.Shiny); err != nil {
		return err
	}

	event.Fire(event.PokemonQuery, p)

	return nil
}

func (p *Pokemon) Insert() error {
	_, err := Connection.Exec(
		"INSERT INTO pokemon (id, species, xp, level, gender, shiny) VALUES (?, ?, ?, ?, ?, ?)",
		p.Id,
		p.Species,
		p.Xp,
		p.Level,
		p.Gender,
		p.Shiny,
	)

	if err == nil {
		event.Fire(event.PokemonInsert, p)
	}

	return err
}

func (p *Pokemon) Delete() error {
	_, err := Connection.Exec("DELETE FROM pokemon WHERE id=?", p.Id)

	if err == nil {
		event.Fire(event.PokemonDelete, p)
	}

	return err
}
