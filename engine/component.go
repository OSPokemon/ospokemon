package engine

import (
	"time"
)

type Components map[string]Component

type Component interface {
	Id() string
	Update(*Universe, *Entity, time.Duration)
}

func (c *Components) AddComponent(comp Component) {
	(*c)[comp.Id()] = comp
}

func (c *Components) Component(key string) Component {
	return (*c)[key]
}
