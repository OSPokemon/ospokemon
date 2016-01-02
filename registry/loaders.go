package registry

import (
	"github.com/ospokemon/ospokemon/world"
)

type LoaderFunc func(int)
type UnloaderFunc func(world.Entity)

var Loaders = make(map[string]LoaderFunc)

var Unloaders = make(map[string]UnloaderFunc)
