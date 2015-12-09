package connection

import (
	"encoding/json"
	"strconv"
)

func UpdateConnections(base map[string]interface{}, cviews map[string]interface{}) {

	for _, client := range Clients {
		view := make(map[string]interface{})
		view["world"] = viewcopier(base).copy()

		controlViews := make(map[string]interface{})
		for _, id := range client.Entities {
			tag := strconv.Itoa(id)
			controlViews[tag] = cviews[tag]
		}
		view["control"] = controlViews

		json, _ := json.Marshal(view)
		message := string(json)

		client.Send <- message
	}
}
