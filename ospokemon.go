package main

import (
	"flag"
	"github.com/ospokemon/ospokemon/connection"
	"github.com/ospokemon/ospokemon/data"
	"github.com/ospokemon/ospokemon/world/update"
	"net/http"
	"time"
)

func main() {
	tickSize := 400
	var port, dir string

	flag.StringVar(&port, "port", "8080", "Port number to open the server on")
	flag.StringVar(&dir, "dir", ".", "A system path to use as web server root")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(dir+"/public/")))
	http.Handle("/connect", connection.ConnectHandler)
	http.Handle("/login", connection.LoginHandler)

	data.Connect(dir + "/db.sqlite")
	go Loop(time.Duration(tickSize) * time.Millisecond)

	http.ListenAndServe(":"+port, nil)
}

func Loop(d time.Duration) {
	for now := range time.Tick(d) {
		view := update.UpdateWorld(now)
		go connection.Update(view, now)
	}
}
