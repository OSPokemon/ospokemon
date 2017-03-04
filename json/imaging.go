package json

import (
	"github.com/ospokemon/ospokemon/game"
)

func Imaging(i *game.Imaging) (string, map[string]interface{}) {
	return "imaging", map[string]interface{}{
		"image":      i.Image,
		"animations": i.Animations,
	}
}
