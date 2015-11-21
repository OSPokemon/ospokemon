package data

import (
	"github.com/ospokemon/ospokemon/world"
	"log"
)

type animationStore byte
type graphicsStore byte

var AnimationStore animationStore
var Animations = make(map[int]map[world.AnimationType]string)

var GraphicsStore graphicsStore

func (a *animationStore) Load(id int) map[world.AnimationType]string {
	if Animations[id] != nil {
		return Animations[id]
	}

	rows, err := Connection.Query("SELECT animation_type, animation FROM animations WHERE id=?", id)
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

	Animations[id] = animations
	return animations
}

func (g *graphicsStore) New(id int) *world.Graphics {
	animations := AnimationStore.Load(id)
	graphics := &world.Graphics{
		Animations: animations,
	}

	return graphics
}
