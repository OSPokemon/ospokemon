package server

import (
	"encoding/json"
	"github.com/ospokemon/ospokemon/objects"
	"net/http"
	"strconv"
	"strings"
)

var TrainerHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	sessionId := readSessionId(r)

	if sessionId < 1 {
		w.WriteHeader(500)
		w.Write([]byte("Session missing"))
	}

	session := Sessions[sessionId]
	account := Accounts[session.Username]

	uriPart := strings.TrimSuffix(strings.TrimPrefix(r.RequestURI, "/trainer/"), ".json")

	id, _ := strconv.ParseInt(uriPart, 10, 0)

	accountTrainer := false
	for _, id2 := range account.TrainerIds {
		accountTrainer = accountTrainer || int(id) == id2
	}
	if !accountTrainer {
		w.WriteHeader(500)
		w.Write([]byte("Disallowed"))
	}

	trainer := objects.GetTrainer(int(id))

	resp := make(map[string]interface{})

	resp["Name"] = trainer.Name()
	resp["Map"] = trainer.Map()
	resp["Pokemon"] = trainer.Pokemon()
	resp["Money"] = trainer.Money()
	resp["Portrait"] = trainer.Graphics()["portrait"]

	trainerJson, _ := json.Marshal(resp)

	w.Write(trainerJson)
})
