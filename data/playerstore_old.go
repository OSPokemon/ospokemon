package data

import (
// "fmt"
// "github.com/ospokemon/ospokemon/api-go"
// "github.com/ospokemon/ospokemon/world"
)

// var Players = make(map[int]Player)
// var PlayersByName = make(map[string]int)

// type Player interface {
// 	ospokemon.Trainer
// 	world.Entity
// 	Password() string
// 	SetPassword(password string)
// }

// type GameServerTrainer struct {
// 	ospokemon.BasicTrainer
// 	position world.Position
// 	size     world.Size
// 	health   int
// 	password string
// 	action   *world.Action
// 	effects  []world.Effect
// }

// func (t *GameServerTrainer) Tag() string {
// 	return fmt.Sprintf("player%i", t.Id())
// }

// func (t *GameServerTrainer) Image() string {
// 	return fmt.Sprintf("trainer/%i.png", t.Class())
// }

// func (t *GameServerTrainer) Position() *world.Position {
// 	return &t.position
// }

// func (t *GameServerTrainer) Size() *world.Size {
// 	return &t.size
// }

// func (t *GameServerTrainer) Solid() bool {
// 	return true
// }

// func (t *GameServerTrainer) Action() *world.Action {
// 	return t.action
// }

// func (t *GameServerTrainer) SetAction(action *world.Action) {
// 	t.action = action
// }

// func (t *GameServerTrainer) Effects() []world.Effect {
// 	return t.effects
// }

// func (t *GameServerTrainer) SetEffects(effects []world.Effect) {
// 	t.effects = effects
// }

// func (t *GameServerTrainer) Health() int {
// 	return t.health
// }

// func (t *GameServerTrainer) SetHealth(health int) {
// 	if health > 100 {
// 		t.health = 100
// 	} else if health < 0 {
// 		t.health = 0
// 	} else {
// 		t.health = health
// 	}
// }

// func (t *GameServerTrainer) Password() string {
// 	return t.password
// }

// func (t *GameServerTrainer) SetPassword(password string) {
// 	t.password = password
// }

// func LoadPlayer(username string) Player {
// 	if PlayersByName[username] > 0 {
// 		return Players[PlayersByName[username]]
// 	}

// 	row := Connection.QueryRow("SELECT * FROM players WHERE name=?", username)
// 	player := &GameServerTrainer{
// 		size: world.Size{64, 64},
// 	}

// 	var id, class, health int
// 	var name, password string
// 	err := row.Scan(&id, &name, &class, &health, &player.position.X, &player.position.Y, &password)

// 	if err != nil {
// 		return nil
// 	}

// 	player.SetId(id)
// 	player.SetName(name)
// 	player.SetClass(class)
// 	player.SetHealth(health)
// 	player.SetPassword(password)
// 	Players[player.Id()] = player
// 	PlayersByName[player.Name()] = player.Id()

// 	return player
// }

// func UnloadPlayer(username string) {

// }
