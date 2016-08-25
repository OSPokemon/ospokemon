package server

import (
	"github.com/ospokemon/ospokemon/util"
	"net/http"
	"time"
)

func Launch() {
	routes()
	go pollSessionExpiry()
	e := http.ListenAndServe(":"+util.Opt("port"), nil)
	util.Log.Error(e)
}

func pollSessionExpiry() {
	for now := range time.Tick(1 * time.Second) {
		for sessionId, session := range Sessions {
			if session.Expire.Before(now) {
				util.Event.Fire(EVNT_SessionExpire, sessionId, session)
			}
		}
	}
}
