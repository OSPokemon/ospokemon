package main

import (
	"flag"
)

var databaseFile, path, port string
var debugMode bool

func flags() {
	flag.StringVar(&databaseFile, "dbfile", "./db.sqlite", "database file")
	flag.StringVar(&path, "path", "./public", "system path to server root")
	flag.StringVar(&port, "port", "8080", "port to open the server on")
	flag.BoolVar(&debugMode, "debug", false, "enable cli logging at DEBUG level")
	flag.Parse()
}
