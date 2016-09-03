package server

import (
	"github.com/ospokemon/ospokemon/util"
	"time"
)

func PollSessionExpire() {
	for now := range time.Tick(1 * time.Second) {
		for sessionId, session := range Sessions {
			if session.Expire.Before(now) {
				util.Event.Fire(EVNT_SessionExpire, sessionId, session)
			}
		}
	}
}
