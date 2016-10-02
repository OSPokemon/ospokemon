package engine

type Script func(*Universe, *Entity, map[string]string)

var Scripts = make(map[uint]Script)
