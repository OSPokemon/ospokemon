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
	http.Handle("/all-classes.json", server.AllClassesHandler)
	http.Handle("/all-species.json", server.AllSpeciesHandler)
	http.Handle("/trainer/", server.TrainerHandler)
	http.Handle("/choose-trainer.go", server.ChooseTrainerHandler)
	http.Handle("/create-trainer.go", server.CreateTrainerHandler)
	http.Handle("/change-password.go", server.ChangePasswordHandler)
	http.Handle("/websocket.go", server.WebsocketHandler)
}
