package data

import (
	"github.com/ospokemon/ospokemon/world"
	"log"
)

type playerStore byte

var PlayerStore playerStore
var Players = make(map[string]*Player)

func (p *playerStore) FetchPassword(name string) string {
	row := Connection.QueryRow("SELECT password FROM players WHERE name=?", name)
	var password string

	err := row.Scan(&password)

	if err != nil {
		return ""
	}

	return password
}

func (p *playerStore) Load(name string) *Player {
	if Players[name] != nil {
		return Players[name]
	}

	row := Connection.QueryRow("SELECT id, name, class, health, maxhealth, x, y FROM players WHERE name=?", name)
	player := &Player{
		SPEED: 25,
		PHYSICS: &world.Physics{
			Position: world.Position{},
			Size:     world.Size{32, 32},
			Solid:    true,
		},
	}

	err := row.Scan(&player.BasicTrainer.ID, &player.NAME, &player.CLASS, &player.HEALTH, &player.MAXHEALTH, &player.Physics().Position.X, &player.Physics().Position.Y)

	if err != nil {
		log.Fatal(err)
	}

	Players[name] = player
	return player
}
