package ospokemon

import (
	"errors"
)

type Scripter struct {
	Script string
	Data   map[string]interface{}
}

func MakeScripter() *Scripter {
	return &Scripter{
		Data: make(map[string]interface{}),
	}
}

func (s *Scripter) Run(e *Entity) error {
	if Scripts[s.Script] == nil {
		return errors.New("ospokemon: missing script")
	}
	return Scripts[s.Script](e, s.Data)
}

var Scripts = make(map[string]func(*Entity, map[string]interface{}) error)
