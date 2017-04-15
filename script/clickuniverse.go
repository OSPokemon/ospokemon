package script

import (
	"errors"
	"ospokemon.com"
	"ospokemon.com/space"
)

func ClickUniverse(e *ospokemon.Entity, data map[string]interface{}) error {
	movement := e.GetMovement()
	if movement == nil {
		return errors.New("clickuniverse: movement missing")
	}

	point := space.Point{}
	if pointx, ok := data["x"].(float64); ok {
		point.X = pointx
	} else {
		return errors.New("clickuniverse: \"x\" missing")
	}
	if pointy, ok := data["y"].(float64); ok {
		point.Y = pointy
	} else {
		return errors.New("clickuniverse: \"y\" missing")
	}

	movement.Up = false
	movement.Down = false
	movement.Left = false
	movement.Right = false
	movement.Target = &point

	return nil
}
