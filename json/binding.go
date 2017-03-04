package json

import (
	"github.com/ospokemon/ospokemon/game"
)

func Binding(b *game.Binding, expand bool) (string, map[string]interface{}) {
	data := map[string]interface{}{
		"key": b.Key,
	}

	if expand {
		for _, part := range b.Parts {
			if imaging, ok := part.(*game.Imaging); ok {
				key, partData := Imaging(imaging)
				data[key] = partData
			} else if walk, ok := part.(game.Walk); ok {
				data["walk"] = walk
			} else if menu, ok := part.(game.Menu); ok {
				data["menu"] = menu
			}
		}
	}

	return "binding", data
}
