package data

// import (
// 	log "github.com/Sirupsen/logrus"
// 	"github.com/ospokemon/ospokemon/world"
// )

// type playerStore byte

// var PlayerStore playerStore
// var Players = make(map[string]*Player)

// func (p *playerStore) Load(name string) *Player {
// 	if Players[name] != nil {
// 		return Players[name]
// 	}

// 	row := Connection.QueryRow("SELECT id, name, class, health, maxhealth, x, y FROM players WHERE name=?", name)
// 	player := &Player{
// 		STATS: map[string]world.Stat{
// 			"health": &PlayerStat{100, 100, 100},
// 			"speed":  &PlayerStat{25, 25, 25},
// 		},
// 		PHYSICS: &world.Physics{
// 			Point: world.Point{},
// 			Size:     world.Size{64, 64},
// 			Solid:    true,
// 		},
// 	}

// 	var health, maxhealth int

// 	err := row.Scan(&player.BasicTrainer.ID, &player.NAME, &player.CLASS, &health, &maxhealth, &player.Physics().Point.X, &player.Physics().Point.Y)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	player.Stats()["health"].SetValue(health)
// 	player.Stats()["health"].SetMaxValue(maxhealth)
// 	player.Stats()["health"].SetBaseMaxValue(maxhealth)

// 	Players[name] = player
// 	return player
// }
