package persistence

import (
	"errors"

	"ospokemon.com"
	"ztaylor.me/log"
)

func DialogsSelect(universe *ospokemon.Universe) (map[uint]*ospokemon.Dialog, error) {
	rows, err := Connection.Query(
		"SELECT entity, id, parent, lead, text FROM dialogs WHERE universe=?",
		universe.Id,
	)
	if err != nil {
		return nil, errors.New("dialogsselect: " + err.Error())
	}

	// dialogs

	dialogs := make(map[uint]map[uint]*ospokemon.Dialog)
	for rows.Next() {
		var entitybuff uint
		dialog := ospokemon.MakeDialog()
		if err = rows.Scan(&entitybuff, &dialog.Id, &dialog.Parent, &dialog.Lead, &dialog.Text); err != nil {
			return nil, err
		}

		if dialogs[entitybuff] == nil {
			dialogs[entitybuff] = make(map[uint]*ospokemon.Dialog)
		}
		dialogs[entitybuff][dialog.Id] = dialog
	}
	rows.Close()

	// dialogs items tests

	rows, err = Connection.Query(
		"SELECT entity, dialog, item, amount FROM dialogs_items_tests WHERE universe=?",
		universe.Id,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		dialogItemTest := &ospokemon.DialogItemTest{}
		var entitybuff, dialogbuff uint

		err = rows.Scan(&entitybuff, &dialogbuff, &dialogItemTest.Item, &dialogItemTest.Amount)
		if err != nil {
			return nil, err
		}

		dialogs[entitybuff][dialogbuff].Tests = append(dialogs[entitybuff][dialogbuff].Tests, dialogItemTest)
	}
	rows.Close()

	// dialogs_scripts

	rows, err = Connection.Query(
		"SELECT entity, dialog, script FROM dialogs_scripts WHERE universe=?",
		universe.Id,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		scripter := ospokemon.MakeScripter()
		var entitybuff, dialogbuff uint

		err = rows.Scan(&entitybuff, &dialogbuff, &scripter.Script)
		if err != nil {
			return nil, err
		}

		dialogs[entitybuff][dialogbuff].Scripts[scripter.Script] = scripter
	}
	rows.Close()

	// dialogs_scripts_data

	rows, err = Connection.Query(
		"SELECT entity, dialog, script, key, value FROM dialogs_scripts_data WHERE universe=?",
		universe.Id,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var entitybuff, dialogbuff uint
		var scriptbuff, keybuff, valuebuff string
		if err = rows.Scan(&entitybuff, &dialogbuff, &scriptbuff, &keybuff, &valuebuff); err != nil {
			return nil, err
		}
		dialogs[entitybuff][dialogbuff].Scripts[scriptbuff].Data[keybuff] = valuebuff
	}
	rows.Close()

	// dialogs compile linked lists

	for _, entityDialogs := range dialogs {
		for _, dialog := range entityDialogs {
			if dialog.Id != 0 {
				parentDialog := entityDialogs[dialog.Parent]
				parentDialog.Choices = append(parentDialog.Choices, dialog)
			}
		}
	}

	// compile down to return dialog#0 for all entities

	compiledDialogs := make(map[uint]*ospokemon.Dialog)
	for entityId, entityDialogs := range dialogs {
		compiledDialogs[entityId] = entityDialogs[0]
	}

	if len(compiledDialogs) > 0 && err == nil {
		log.Add("Universe", universe.Id).Add("Dialogs", len(compiledDialogs)).Debug("dialogs select")
	}

	return compiledDialogs, nil
}
