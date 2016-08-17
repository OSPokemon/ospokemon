package engine

import (
	"time"
)

type Component interface {
	Id() string
	Update(*Universe, *Entity, time.Duration)
}
