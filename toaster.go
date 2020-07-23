package ospokemon

import (
	"taylz.io/types"
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

func (toast *Toast) Json() types.Dict {
	return types.Dict{
		"color":   toast.Color,
		"image":   toast.Image,
		"message": toast.Message,
	}
}

func (toaster *Toaster) Json() types.Dict {
	data := types.Dict{}
	for id, toast := range *toaster {
		data[types.StringInt(id)] = toast.Json()
	}
	return data
}
