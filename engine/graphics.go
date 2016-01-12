package engine

import (
	"github.com/ospokemon/ospokemon/physics"
	"math"
)

type AnimationType string

const (
	ANIMportrait   AnimationType = "portrait"
	ANIMwalk_down  AnimationType = "walkdown"
	ANIMwalk_right AnimationType = "walkright"
	ANIMwalk_up    AnimationType = "walkup"
	ANIMwalk_left  AnimationType = "walkleft"
)

func DirectionAnimation(v physics.Vector) AnimationType {
	if math.Abs(v.DX) > math.Abs(v.DY) {
		if v.DX > 0 {
			return ANIMwalk_right
		} else {
			return ANIMwalk_left
		}
	} else {
		if v.DY > 0 {
			return ANIMwalk_down
		} else {
			return ANIMwalk_up
		}
	}
}
