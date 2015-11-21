package data

import (
// "fmt"
// "src.ospokemon.org/gameserver/world"
// "src.ospokemon.org/ospokemon"
)

// var Pokemon = make(map[int]PokemonEntity)

// type PokemonEntity interface {
// 	ospokemon.Pokemon
// 	world.Entity
// }

// type GameServerPokemon struct {
// 	ospokemon.BasicPokemon
// 	position world.Position
// 	size     world.Size
// 	action   *world.Action
// 	effects  []world.Effect
// }

// func (p *GameServerPokemon) Tag() string {
// 	return fmt.Sprintf("pokemon%i", p.Id())
// }

// func (p *GameServerPokemon) Image() string {
// 	return fmt.Sprintf("trainer/%i.png", p.Id())
// }

// func (p *GameServerPokemon) Position() *world.Position {
// 	return &p.position
// }

// func (p *GameServerPokemon) Size() *world.Size {
// 	return &p.size
// }

// func (p *GameServerPokemon) Solid() bool {
// 	return true
// }

// func (p *GameServerPokemon) Action() *world.Action {
// 	return p.action
// }

// func (p *GameServerPokemon) SetAction(action *world.Action) {
// 	p.action = action
// }

// func (p *GameServerPokemon) Effects() []world.Effect {
// 	return p.effects
// }

// func (p *GameServerPokemon) SetEffects(effects []world.Effect) {
// 	p.effects = effects
// }

// func (p *GameServerPokemon) Health() int {
// 	stats := p.Stats()
// 	healthstat := stats["health"]
// 	return healthstat.Value()
// }

// func (p *GameServerPokemon) SetHealth(health int) {
// 	// ( { [IV+2*Base Stat+([EVs]/4)+100] * Level } / 100 )+10
// 	stats := p.Stats()
// 	healthstat := stats["health"]

// 	if health < 0 {
// 		healthstat.SetValue(0)
// 	}

// 	maxhealth := ((healthstat.EffortValue() / 4) + 100)

// 	if health > maxhealth {
// 		healthstat.SetValue(maxhealth)
// 	}
// }
