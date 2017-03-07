package game

type Script func(*Entity, map[string]interface{}) error

var Scripts = make(map[string]Script)
