package ospokemon

import "taylz.io/types"

const PARTdialog = "dialog"

type Dialog struct {
	Id      uint
	Parent  uint
	Lead    string
	Text    string
	Tests   []DialogTest
	Choices []*Dialog
	Scripts map[string]*Scripter
}

func MakeDialog() *Dialog {
	return &Dialog{
		Tests:   make([]DialogTest, 0),
		Choices: make([]*Dialog, 0),
		Scripts: make(map[string]*Scripter),
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

func (dialog *Dialog) Json() types.Dict {
	json := types.Dict{
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
