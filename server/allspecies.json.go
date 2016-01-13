package server

import (
	"encoding/json"
	"github.com/ospokemon/ospokemon/objects"
	"net/http"
)

var AllSpeciesHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	resp := make([]interface{}, 0)

	for _, species := range objects.GetAllSpecies() {
		speciesJson := make(map[string]interface{})
		speciesJson["Id"] = species.Id()
		speciesJson["Name"] = species.Name()
		speciesJson["Graphic"] = species.GRAPHICS["portrait"]
		speciesJson["Stats"] = species.Stats()
		resp = append(resp, speciesJson)
	}

	allSpeciesJson, _ := json.Marshal(resp)

	w.Write(allSpeciesJson)
})
