package server

import (
	"github.com/ospokemon/ospokemon/objects"
	"net/http"
	"strconv"
)

var CreateTrainerHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(404)
		return
	}

	sessionId := readSessionId(r)
	if sessionId < 1 {
		w.WriteHeader(500)
		w.Write([]byte("SessionId missing"))
	}

	session := Sessions[sessionId]
	if session == nil {
		w.WriteHeader(500)
		w.Write([]byte("Session missing"))
	}

	name := r.FormValue("name")
	class, err := strconv.ParseInt(r.FormValue("class"), 10, 0)
	if err != nil {
		http.Redirect(w, r, "/create-trainer", http.StatusMovedPermanently)
		return
	}

	trainer, err := objects.CreateTrainer(name, int(class))
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	speciesId, err := strconv.ParseInt(r.FormValue("species"), 10, 0)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	pokemon, err := objects.MakePokemon(int(speciesId))
	pokemon.ORIGINALTRAINER = trainer.ID
	err = objects.SavePokemon(pokemon)

	http.Redirect(w, r, "/login", http.StatusMovedPermanently)
})
