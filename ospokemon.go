package main

import (
	"flag"
	log "github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/loader"
	_ "github.com/ospokemon/ospokemon/objects"
	"github.com/ospokemon/ospokemon/server"
	"github.com/ospokemon/ospokemon/update"
	"net/http"
	"time"
)

var databaseFile, path, port string
var debugMode = false

func main() {
	readFlags()
	configureLoader()

	if debugMode {
		log.SetLevel(log.DebugLevel)
	}

	tickSize := 250
	go Loop(time.Duration(tickSize) * time.Millisecond)

	http.Handle("/", http.FileServer(http.Dir(path)))
	http.Handle("/connect", server.WebsocketHandler)
	http.Handle("/login", server.LoginHandler)
	http.ListenAndServe(":"+port, nil)
}

func readFlags() {
	flag.StringVar(&databaseFile, "db", "db.sqlite", "database file")
	flag.StringVar(&path, "path", "./public/", "system path to server root")
	flag.StringVar(&port, "port", "8080", "port to open the server on")
	flag.BoolVar(&debugMode, "debug", false, "enable cli logging at DEBUG level")
	flag.Parse()
}

func configureLoader() {
	server.LoginAccount = loader.LoginAccount
	server.ConnectClient = loader.ConnectClient
	server.DisconnectClient = loader.DisconnectClient
	server.ReceiveMessage = update.ReceiveMessage

	loader.Connect(databaseFile)
	loader.LoadAllSpells()
}

func Loop(d time.Duration) {
	for now := range time.Tick(d) {
		view, cview := update.UpdateWorld(now)
		server.UpdateClients(view, cview)
	}
}
