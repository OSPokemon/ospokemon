package server

import (
	"encoding/json"
	"github.com/ospokemon/ospokemon/objects"
	"net/http"
)

var AllClassesHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	resp := make([]interface{}, 0)

	for _, class := range objects.GetAllClasses() {
		classJson := make(map[string]interface{})
		classJson["Id"] = class.Id()
		classJson["Name"] = class.Name()
		classJson["Graphic"] = class.Graphics["portrait"]
		resp = append(resp, classJson)
	}

	allClassesJson, _ := json.Marshal(resp)

	w.Write(allClassesJson)
})
