package world

import (
	"time"
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

type Contour struct {
	Spritesheet string
	X           int
	Y           int
}

type Animation struct {
	Speed  time.Duration
	States []*Contour
}

type Animating struct {
	Clock time.Time
	Frame int
	AnimationType
}

type Graphics struct {
	Current    *Animating
	Animations map[AnimationType]*Animation
}

func (g *Graphics) Update(e Entity, now time.Time) bool {
	timeDiff := now.Sub(g.Current.Clock)
	animation := g.Animations[g.Current.AnimationType]
	frame := int(int64(timeDiff) / int64(animation.Speed))

	if frame != g.Current.Frame && frame < len(animation.States) {
		g.Current.Frame = frame
		return true
	}

	return false
}
