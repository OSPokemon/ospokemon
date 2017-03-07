package game

import (
	"github.com/ospokemon/ospokemon/part"
)

type Dialog struct {
	Id      uint
	Parent  uint
	Lead    string
	Text    string
	Choices []*Dialog
	Script  string
	Data    map[string]interface{}
}

func MakeDialog() *Dialog {
	return &Dialog{
		Choices: make([]*Dialog, 0),
		Data:    make(map[string]interface{}),
	}
}

func (d *Dialog) Next(lead string) *Dialog {
	for _, choice := range d.Choices {
		if choice.Lead == lead {
			return choice
		}
	}

	return nil
}

func (d *Dialog) Part() string {
	return part.Dialog
}
