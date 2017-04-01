package run

import (
	"ospokemon.com/game"
	"ospokemon.com/option"
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
