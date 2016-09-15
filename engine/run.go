package engine

import (
	"github.com/ospokemon/ospokemon/util"
	"time"
)

func Run() {
	d := time.Duration(util.OptInt("sessionlife")) * time.Millisecond

	for range time.Tick(d) {
		for _, universe := range Multiverse {
			universe.Update(d)
		}
	}
}
