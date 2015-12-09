package loader

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/world"
)

var Animations = make(map[string]map[int]map[world.AnimationType]string)

func LoadAnimations(t string, id int) {
	if Animations[t] == nil {
		Animations[t] = make(map[int]map[world.AnimationType]string)
	}
	if Animations[t][id] != nil {
		return
	}

	rows, err := Connection.Query("SELECT animationtype, animation FROM animations WHERE type=? AND id=?", t, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	animations := make(map[world.AnimationType]string)
	var animType int
	var animation string

	for rows.Next() {
		err = rows.Scan(&animType, &animation)

		if err != nil {
			log.Fatal(err)
		}

		animations[world.AnimationType(animType)] = animation
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	Animations[t][id] = animations

	log.WithFields(log.Fields{
		"type":       t,
		"id":         id,
		"animations": animations,
	}).Debug("Animations built")
}
