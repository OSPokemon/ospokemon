package signup

import (
	"ospokemon.com"
)

var menuBindings = map[string]ospokemon.Menu{
	"Enter":  "chat",
	"c":      "player",
	"b":      "itembag",
	"x":      "actions",
	"Escape": "settings",
}

var movementBindings = map[string]ospokemon.Walk{
	"a": "left",
	"s": "down",
	"d": "right",
	"w": "up",
}
