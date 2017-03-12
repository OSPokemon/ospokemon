package game

import (
	"github.com/ospokemon/ospokemon/json"
)

const PARTimaging = "imaging"

type Imaging struct {
	Image      string
	Animations map[string]string
}

func MakeImaging() *Imaging {
	return &Imaging{
		Animations: make(map[string]string),
	}
}

func BuildImaging(sample map[string]string) *Imaging {
	return MakeImaging().ReadAnimations(sample)
}

func (i *Imaging) ReadAnimations(sample map[string]string) *Imaging {
	i.Image = sample["portrait"]
	for k, v := range sample {
		i.Animations[k] = v
	}
	return i
}

func (i *Imaging) Part() string {
	return PARTimaging
}

func (parts Parts) GetImaging() *Imaging {
	imaging, _ := parts[PARTimaging].(*Imaging)
	return imaging
}

func (imaging *Imaging) Json() json.Json {
	return json.Json{
		"image":      imaging.Image,
		"animations": imaging.Animations,
	}
}
