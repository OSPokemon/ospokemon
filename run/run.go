package run

import (
	"time"

	"ospokemon.com"
	"ospokemon.com/server/sessionman"
	"ztaylor.me/cast"
	"ztaylor.me/env"
)

func Run(env env.Provider) {
	d := time.Duration(cast.Int(env.Get("refresh"))) * time.Millisecond
	for range time.Tick(d) {
		for _, u := range ospokemon.Universes.Cache {
			u.Update(d)
		}
		for _, s := range sessionman.Cache {
			s.Frame()
		}
	}
}
