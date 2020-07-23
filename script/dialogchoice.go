package script

import (
	"errors"
	"github.com/ospokemon/ospokemon"
)

func DialogChoice(e *ospokemon.Entity, choice string) error {
	if choice == "" {
		return errors.New("dialogchoice: missing data")
	}

	dialog := e.GetDialog()
	if dialog == nil {
		return errors.New("dialogchoice: dialog missing")
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
