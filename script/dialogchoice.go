package script

import (
	"errors"
	"ospokemon.com"
)

func init() {
	ospokemon.Scripts["dialogchoice"] = DialogChoice
}

func DialogChoice(e *ospokemon.Entity, data map[string]interface{}) error {
	dialog := e.GetDialog()
	if dialog == nil {
		return nil
	}

	choice, _ := data["choice"].(string)
	if choice == "" {
		return nil
	}

	if nextDialog := dialog.Next(choice); nextDialog != nil {
		for _, tester := range nextDialog.Tests {
			if !tester.Test(e.GetPlayer()) {
				return errors.New("dialogchoice: test failure")
			}
		}

		e.AddPart(nextDialog)

		for _, scripter := range nextDialog.Scripts {
			if err := scripter.Run(e); err != nil {
				return errors.New("dialogchoice: " + err.Error())
			}
		}
	} else if len(dialog.Choices) < 1 {
		e.RemovePart(dialog)
	}

	return nil
}
