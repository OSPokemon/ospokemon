package run

import (
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/option"
	"time"
)

func Run() {
	d := time.Duration(option.Int("refresh")) * time.Millisecond

	for range time.Tick(d) {
		for _, universe := range game.Multiverse {
			universe.Update(d)
		}
	}
}
