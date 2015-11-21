package connection

import (
	"encoding/json"
	"github.com/ospokemon/ospokemon/world"
	"log"
	"strconv"
)

func Update(base map[string]*world.View) {
	// json, err := json.Marshal(view)

	// if err != nil {
	// 	log.Printf("connection.Update err: ", err)
	// 	return
	// }

	// message := string(json)

	// for _, client := range Clients {
	// 	client.Send <- message
	// }

	for _, client := range Clients {
		view := make(map[string]interface{})
		view["world"] = copyMap(base)

		controlViews := make(map[string]*world.View)
		for _, id := range client.Entities {
			tag := strconv.Itoa(id)
			controlViews[tag] = base[tag]
		}
		view["control"] = controlViews

		json, _ := json.Marshal(view)
		message := string(json)

		client.Send <- message
	}
}

func ReceiveMessage(name string, message map[string]interface{}) {
	log.Printf("websocket receive from:%s message: %v\n", name, message)
}

func copyMap(src map[string]*world.View) map[string]*world.View {
	dst := make(map[string]*world.View)

	for k, v := range src {
		dst[k] = v
	}

	return dst
}
