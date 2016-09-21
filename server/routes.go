package server

import (
	"net/http"
)

func routes() {
	http.Handle("/", PageHandler)
	http.Handle("/api/signup", SignupHandler)
	http.Handle("/api/login", LoginHandler)
	http.Handle("/api/logout", LogoutHandler)
	http.Handle("/api/player", PlayerHandler)
	http.Handle("/api/websocket", WebsocketHandler)
}
