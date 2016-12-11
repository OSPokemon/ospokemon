package run

import (
	"github.com/ospokemon/ospokemon/option"
	"github.com/ospokemon/ospokemon/save"
	"time"
)

func Run() {
	d := time.Duration(option.Int("refresh")) * time.Millisecond

	for range time.Tick(d) {
		for _, universe := range save.Multiverse {
			universe.Update(d)
		}
	}
}
