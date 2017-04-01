package game

import (
	"ospokemon.com/json"
)

const PARTdialog = "dialog"

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
	return PARTdialog
}

func (parts Parts) GetDialog() *Dialog {
	dialog, _ := parts[PARTdialog].(*Dialog)
	return dialog
}

func (dialog *Dialog) Json() json.Json {
	json := json.Json{
		"id":   dialog.Id,
		"lead": dialog.Lead,
		"text": dialog.Text,
	}
	choices := make([]string, 0)
	for _, choice := range dialog.Choices {
		choices = append(choices, choice.Lead)
	}
	json["choices"] = choices
	return json
}
