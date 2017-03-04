package game

type Script func(*Entity, map[string]string) error

var Scripts = make(map[string]Script)
