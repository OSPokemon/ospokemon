package run

import (
	"ospokemon.com"
	"ospokemon.com/option"
	"ospokemon.com/server/session"
	"time"
)

func Run() {
	d := time.Duration(option.Int("refresh")) * time.Millisecond

	for range time.Tick(d) {
		for _, universe := range ospokemon.Multiverse {
			universe.Update(d)
		}

		for _, s := range session.Sessions {
			s.Frame()
		}
	}
}
