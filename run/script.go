package run

import (
	"github.com/ospokemon/ospokemon/save"
)

type Script func(*save.Universe, *save.Entity, map[string]string)

var Scripts = make(map[uint]Script)
