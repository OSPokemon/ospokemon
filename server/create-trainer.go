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
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}

	if objects.CreateTrainer == nil {
		w.Write([]byte("SHIT"))
	}

	trainer, err := objects.CreateTrainer(session.Username, name, int(class))
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

	pokemonname := r.FormValue("pokemonname")

	pokemon, err := objects.MakePokemon(pokemonname, int(speciesId))
	pokemon.ORIGINALTRAINER = trainer.ID
	err = objects.SavePokemon(pokemon)

	http.Redirect(w, r, "/dashboard", http.StatusMovedPermanently)
})
