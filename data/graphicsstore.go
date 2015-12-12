// package data

// import (
// 	log "github.com/Sirupsen/logrus"
// 	"github.com/ospokemon/ospokemon/world"
// )

// type animationStore byte
// type graphicsStore byte

// var AnimationStore animationStore
// var GraphicsStore graphicsStore

// var Animations = make(map[string]map[int]map[world.AnimationType]string)

// func (a *animationStore) Load(t string, id int) map[world.AnimationType]string {
// 	if Animations[t] == nil {
// 		Animations[t] = make(map[int]map[world.AnimationType]string)
// 	}
// 	if Animations[t][id] != nil {
// 		return Animations[t][id]
// 	}

// 	rows, err := Connection.Query("SELECT animationtype, animation FROM animations WHERE type=? AND id=?", t, id)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	animations := make(map[world.AnimationType]string)
// 	var animType int
// 	var animation string

// 	for rows.Next() {
// 		err = rows.Scan(&animType, &animation)

// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		animations[world.AnimationType(animType)] = animation
// 	}

// 	err = rows.Err()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	Animations[t][id] = animations

// 	log.WithFields(log.Fields{
// 		"type":       t,
// 		"id":         id,
// 		"animations": animations,
// 	}).Debug("Animations provided")

// 	return animations
// }

// func (g *graphicsStore) New(t string, id int) *world.Graphics {
// 	animations := AnimationStore.Load(t, id)
// 	return &world.Graphics{
// 		Portrait:   animations[world.ANIMportrait],
// 		Current:    animations[world.ANIMwalk_down],
// 		Animations: animations,
// 	}
// }
