package json

import (
	"github.com/ospokemon/ospokemon/game"
)

func Dialog(dialog *game.Dialog) map[string]interface{} {
	data := map[string]interface{}{
		"id":   dialog.Id,
		"lead": dialog.Lead,
		"text": dialog.Text,
	}

	choices := make([]string, 0)
	for _, choice := range dialog.Choices {
		choices = append(choices, choice.Lead)
	}
	data["choices"] = choices

	return data
}
