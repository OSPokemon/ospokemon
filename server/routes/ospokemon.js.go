package routes

import (
	"net/http"
	"ospokemon.com/server/routes/ospokemon.js"
)

var OSPokemonJs = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if ospokemonjs.Content == "" {
		ospokemonjs.CreateContent()
	}

	w.Header().Set("Content-Type", "application/javascript")
	w.Write([]byte(ospokemonjs.Content))
})
