package objects

import (
	"github.com/ospokemon/api-go"
	"github.com/ospokemon/ospokemon/engine"
	"github.com/ospokemon/ospokemon/physics"
)

type Pokemon struct {
	ENTITYID  int
	GRAPHIC   string
	COLLISION engine.Collision
	SHAPE     physics.Shape
	ospokemon.BasicPokemon
	ACTION    *engine.Action
	CONTROL   engine.Control
	STATS     map[string]*engine.Stat
	ABILITIES []*engine.Ability
	GRAPHICS  map[engine.AnimationType]string
	EFFECTS   []*engine.Effect
	WALKING   *physics.Point
}

var Pokemons = make(map[int]*Pokemon)

var LoadPokemon func(pokemonId int) (*Pokemon, error)
var MakePokemon func(speciesId int) (*Pokemon, error) // called Make because its not supposed to database
var SavePokemon func(pokemon *Pokemon) error

// Pokemon is an entity

func (p *Pokemon) EntityId() *int {
	return &p.ENTITYID
}

func (p *Pokemon) Graphic() *string {
	return &p.GRAPHIC
}

func (p *Pokemon) Collision() *engine.Collision {
	return &p.COLLISION
}

func (p *Pokemon) Shape() physics.Shape {
	return p.SHAPE
}

func (p *Pokemon) SetShape(shape physics.Shape) {
	p.SHAPE = shape
}

func (p *Pokemon) Action() *engine.Action {
	return p.ACTION
}

func (p *Pokemon) SetAction(action *engine.Action) {
	p.ACTION = action
}

func (p *Pokemon) Control() *engine.Control {
	return &p.CONTROL
}

func (p *Pokemon) Abilities() *[]*engine.Ability {
	return &p.ABILITIES
}

func (p *Pokemon) Stats() map[string]*engine.Stat {
	return p.STATS
}

func (p *Pokemon) Graphics() map[engine.AnimationType]string {
	return p.GRAPHICS
}

func (p *Pokemon) Effects() *[]*engine.Effect {
	return &p.EFFECTS
}

func (p *Pokemon) Walking() *physics.Point {
	return p.WALKING
}

func (p *Pokemon) SetWalking(walking *physics.Point) {
	p.WALKING = walking
}

func GetPokemon(pokemonId int) *Pokemon {
	if Pokemons[pokemonId] == nil {
		pokemon, err := LoadPokemon(pokemonId)

		if err != nil {
			return nil
		}

		Pokemons[pokemonId] = pokemon
	}

	return Pokemons[pokemonId]
}
