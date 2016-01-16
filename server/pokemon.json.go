package server

import (
	"encoding/json"
	"github.com/ospokemon/ospokemon/objects"
	"net/http"
	"strconv"
	"strings"
)

var PokemonHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	sessionId := readSessionId(r)

	if sessionId < 1 {
		w.WriteHeader(500)
		w.Write([]byte("Session missing"))
	}

	session := Sessions[sessionId]
	account := Accounts[session.Username]

	uriPart := strings.TrimSuffix(strings.TrimPrefix(r.RequestURI, "/pokemon/"), ".json")

	id, _ := strconv.ParseInt(uriPart, 10, 0)

	accountPokemon := false
	for _, trainerId := range account.TrainerIds {
		trainer := objects.GetTrainer(trainerId)
		for _, pokemonId := range trainer.Pokemon() {
			accountPokemon = accountPokemon || pokemonId == int(id)
		}
	}
	if !accountPokemon {
		w.WriteHeader(500)
		w.Write([]byte("Disallowed"))
	}

	pokemon := objects.GetPokemon(int(id))

	resp := make(map[string]interface{})

	resp["Name"] = pokemon.Name()
	resp["Level"] = pokemon.Level()
	resp["Experience"] = pokemon.Experience()
	resp["Portrait"] = pokemon.Graphics()["portrait"]

	pokemonJson, _ := json.Marshal(resp)

	w.Write(pokemonJson)
})
