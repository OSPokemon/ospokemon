package game

import (
	"github.com/ospokemon/ospokemon/json"
)

const PARTtoaster = "toaster"

type Toast struct {
	Color   string
	Image   string
	Message string
}

type Toaster []*Toast

func MakeToaster() *Toaster {
	toaster := make(Toaster, 0)
	return &toaster
}

func (toaster *Toaster) Part() string {
	return PARTtoaster
}

func (parts Parts) GetToaster() *Toaster {
	toaster, _ := parts[PARTtoaster].(*Toaster)
	return toaster
}

func (toaster *Toaster) Add(toast *Toast) {
	*toaster = append(*toaster, toast)
}

func (toaster *Toaster) AddError(err error) {
	toaster.Add(&Toast{
		Color:   "red",
		Message: err.Error(),
	})
}

func (toaster *Toaster) Clear() {
	*toaster = *MakeToaster()
}

func (toast *Toast) Json() json.Json {
	return json.Json{
		"color":   toast.Color,
		"image":   toast.Image,
		"message": toast.Message,
	}
}

func (toaster *Toaster) Json() json.Json {
	data := make(json.Json)
	for id, toast := range *toaster {
		data[json.StringInt(id)] = toast.Json()
	}
	return data
}
