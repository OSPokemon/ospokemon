package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/db"
	"github.com/ospokemon/ospokemon/engine"
	"github.com/ospokemon/ospokemon/server"
	"github.com/ospokemon/ospokemon/snapshot"
	"net/http"
	"time"
)

func main() {
	flags()
	wiring()
	routes()

	if debugMode {
		log.SetLevel(log.DebugLevel)
	}

	go Loop(time.Duration(250) * time.Millisecond)

	db.Connect(databaseFile)
	http.ListenAndServe(":"+port, nil)
}

func Loop(d time.Duration) {
	for now := range time.Tick(d) {
		for _, m := range engine.Maps {
			engine.UpdateMap(m, now)
			m.Lock()
			view, cview := snapshot.Make(m, now)
			server.PushSnapshot(m.Name, m.Clients, view, cview)
			m.Unlock()
		}
	}
}
