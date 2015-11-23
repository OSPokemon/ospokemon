package data

import (
	"github.com/ospokemon/ospokemon/world"
	"log"
)

type animationStore byte
type graphicsStore byte

var AnimationStore animationStore
var GraphicsStore graphicsStore

var Animations = make(map[string]map[int]map[world.AnimationType]string)

func (a *animationStore) Load(t string, id int) map[world.AnimationType]string {
	if Animations[t] == nil {
		Animations[t] = make(map[int]map[world.AnimationType]string)
	}
	if Animations[t][id] != nil {
		return Animations[t][id]
	}

	rows, err := Connection.Query("SELECT animationtype, animation FROM animations WHERE type=? AND id=?", t, id)
	defer rows.Close()

	if err != nil {
		log.Fatal(err)
	}

	animations := make(map[world.AnimationType]string)
	var animTypeInt int
	var animType world.AnimationType
	var animation string

	for rows.Next() {
		err = rows.Scan(&animTypeInt, &animation)

		if err != nil {
			log.Fatal(err)
		}

		animType = world.AnimationType(animTypeInt)
		animations[animType] = animation
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	Animations[t][id] = animations

	log.Printf("Animations provided for %s#%d: %v", t, id, animations)

	return animations
}

func (g *graphicsStore) New(t string, id int) *world.Graphics {
	animations := AnimationStore.Load(t, id)
	return &world.Graphics{
		Current:    animations[world.ANIMwalk_down],
		Animations: animations,
	}
}
