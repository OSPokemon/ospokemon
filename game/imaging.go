package game

import (
	"github.com/ospokemon/ospokemon/part"
)

type Imaging struct {
	Image      string
	Animations map[string]string
}

func MakeImaging() *Imaging {
	return &Imaging{
		Animations: make(map[string]string),
	}
}

func (i *Imaging) ReadAnimations(sample map[string]string) *Imaging {
	i.Image = sample["portrait"]
	for k, v := range sample {
		i.Animations[k] = v
	}
	return i
}

func (i *Imaging) Part() string {
	return part.Imaging
}
