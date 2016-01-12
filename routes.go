package main

import (
	"github.com/ospokemon/ospokemon/server"
	"net/http"
)

func routes() {
	http.Handle("/", http.FileServer(http.Dir(path)))
	http.Handle("/signup.go", server.SignupHandler)
	http.Handle("/login.go", server.LoginHandler)
	http.Handle("/logout.go", server.LogoutHandler)
	http.Handle("/account.json", server.AccountHandler)
	http.Handle("/choose-profile.go", server.ChooseProfileHandler)
	http.Handle("/change-password.go", server.ChangePasswordHandler)
	http.Handle("/websocket.go", server.WebsocketHandler)
}
