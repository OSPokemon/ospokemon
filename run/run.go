package run

import (
	"github.com/ospokemon/ospokemon"
	"github.com/ospokemon/ospokemon/server/sessionman"
	"taylz.io/env"
	"taylz.io/types"
)

func Run(env env.Service) {
	d := types.Duration(types.IntString(env["refresh"])) * types.Millisecond
	for range types.NewChanTick(d) {
		for _, u := range ospokemon.Universes.Cache {
			u.Update(d)
		}
		for _, s := range sessionman.Cache {
			s.Frame()
		}
	}
}
