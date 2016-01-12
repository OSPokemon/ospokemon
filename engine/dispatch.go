package engine

import (
	// log "github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/physics"
	// "github.com/ospokemon/ospokemon/server"
	// "github.com/ospokemon/ospokemon/world"
)

func ReceiveMessage(username string, message map[string]interface{}) {

}

func parsePoint(src interface{}) physics.Point {
	coords := src.(map[string]interface{})
	point := physics.Point{}
	point.X = coords["x"].(float64)
	point.Y = coords["y"].(float64)
	return point
}
