package json

import (
	"github.com/ospokemon/ospokemon/game"
)

func Bindings(b game.Bindings) (string, map[string]interface{}) {
	data := make(map[string]interface{})

	for key, binding := range b {
		_, bindingData := Binding(binding, true)
		data[key] = bindingData
	}

	return "bindings", data
}
