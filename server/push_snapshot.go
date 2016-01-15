package server

import (
	"encoding/json"
	"strconv"
)

func PushSnapshot(name string, clientIds []int, base map[string]interface{}, cviews map[string]interface{}) {

	for _, clientId := range clientIds {
		client := Clients[clientId]
		view := make(map[string]interface{})
		view["world"] = copyView(base)

		controlViews := make(map[string]interface{})
		for _, id := range client.Entities {
			tag := strconv.Itoa(id)
			controlViews[tag] = cviews[tag]
		}
		view["control"] = controlViews

		json, _ := json.Marshal(view)
		message := string(json)

		client.Send(message)
	}
}

func copyView(src map[string]interface{}) map[string]interface{} {
	dst := make(map[string]interface{})

	for k, v := range src {
		dst[k] = v
	}

	return dst
}
