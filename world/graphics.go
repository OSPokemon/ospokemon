package world

import (
	"math"
)

type AnimationType uint8

const (
	ANIMwalk_down AnimationType = iota
	ANIMwalk_right
	ANIMwalk_up
	ANIMwalk_left
	ANIMcast
	ANIMstun
)

type Graphics struct {
	Current    string
	Animations map[AnimationType]string
}

func (v *Vector) AnimationType() AnimationType {
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
