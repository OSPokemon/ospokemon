package engine

type Script func(*Universe, *Entity, map[string]interface{})

var Scripts = make(map[string]Script)
