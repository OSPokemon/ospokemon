package main

import (
	"flag"
	log "github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/connection"
	"github.com/ospokemon/ospokemon/loader"
	_ "github.com/ospokemon/ospokemon/objects"
	"github.com/ospokemon/ospokemon/update"
	"net/http"
	"time"
)

func main() {
	tickSize := 250
	var port, dir, logLevel string

	flag.StringVar(&port, "port", "8080", "Port number to open the server on")
	flag.StringVar(&dir, "dir", ".", "A system path to use as web server root")
	flag.StringVar(&logLevel, "log", ".", "Log level to set as minimum")
	flag.Parse()

	if logLevel == "DEBUG" {
		log.SetLevel(log.DebugLevel)
	}

	http.Handle("/", http.FileServer(http.Dir(dir+"/public/")))
	http.Handle("/connect", connection.ConnectHandler)
	http.Handle("/login", connection.LoginHandler)

	loader.Connect(dir + "/db.sqlite")
	loader.LoadAllSpells()

	go Loop(time.Duration(tickSize) * time.Millisecond)

	http.ListenAndServe(":"+port, nil)
}

func Loop(d time.Duration) {
	for now := range time.Tick(d) {
		view, cview := update.UpdateWorld(now)
		connection.UpdateConnections(view, cview)
	}
}
