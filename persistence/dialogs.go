package persistence

import (
	"strings"

	"github.com/pkg/errors"
	"ospokemon.com"
	"ztaylor.me/cast"
	"ztaylor.me/log"
)

func DialogsSelect(universeID uint) (map[uint]*ospokemon.Dialog, error) {
	rows, err := Connection.Query(
		"SELECT entity, id, parent, lead, text FROM dialogs WHERE universe=?",
		universeID,
	)
	if err != nil {
		return nil, errors.Wrap(err, "dialogs.select")
	}

	// dialogs

	dialogs := make(map[uint]map[uint]*ospokemon.Dialog)
	for rows.Next() {
		var entitybuff uint
		dialog := ospokemon.MakeDialog()
		if err = rows.Scan(&entitybuff, &dialog.Id, &dialog.Parent, &dialog.Lead, &dialog.Text); err != nil {
			return nil, errors.Wrap(err, "dialogs.scan")
		}

		if dialogs[entitybuff] == nil {
			dialogs[entitybuff] = make(map[uint]*ospokemon.Dialog)
		}
		dialogs[entitybuff][dialog.Id] = dialog
	}
	rows.Close()

	// dialogs items tests item:quanity (i:q)

	rows, err = Connection.Query(
		"SELECT entity, dialog, data FROM dialogs_tests WHERE universe=? AND test='i:q'",
		universeID,
	)
	if err != nil {
		return nil, errors.Wrap(err, "dialogs_tests.select")
	}

	for rows.Next() {
		dialogItemTest := &ospokemon.DialogItemTest{}
		var dialogTestData string
		var entitybuff, dialogbuff uint

		err = rows.Scan(&entitybuff, &dialogbuff, &dialogTestData)
		if err != nil {
			return nil, errors.Wrap(err, "dialogs_tests.scan")
		}

		if parts := strings.Split(dialogTestData, ":"); len(parts) != 2 {
			log.Add("Parts", parts).Warn("dialogs_tests: i:q data invalid")
		} else if i := cast.Int(parts[0]); i < 1 {
			log.Add("I", parts[0]).Warn("dialogs_tests: i:q i invalid")
		} else if q := cast.Int(parts[1]); q < 1 {
			log.Add("I", parts[0]).Warn("dialogs_tests: i:q q invalid")
		} else {
			dialogItemTest.Item = uint(i)
			dialogItemTest.Amount = q
			dialogs[entitybuff][dialogbuff].Tests = append(
				dialogs[entitybuff][dialogbuff].Tests,
				dialogItemTest,
			)
		}
	}
	rows.Close()

	// dialogs_scripts

	rows, err = Connection.Query(
		"SELECT entity, dialog, script FROM dialogs_scripts WHERE universe=?",
		universeID,
	)
	if err != nil {
		return nil, errors.Wrap(err, "dialogs_scripts.select")
	}

	for rows.Next() {
		scripter := ospokemon.MakeScripter()
		var entitybuff, dialogbuff uint

		err = rows.Scan(&entitybuff, &dialogbuff, &scripter.Script)
		if err != nil {
			return nil, errors.Wrap(err, "dialogs_scripts.scan")
		}

		dialogs[entitybuff][dialogbuff].Scripts[scripter.Script] = scripter
	}
	rows.Close()

	// dialogs_scripts_data

	rows, err = Connection.Query(
		"SELECT entity, dialog, script, `key`, value FROM dialogs_scripts_data WHERE universe=?",
		universeID,
	)
	if err != nil {
		return nil, errors.Wrap(err, "dialogs_scripts_data.select")
	}

	for rows.Next() {
		var entitybuff, dialogbuff uint
		var scriptbuff, keybuff, valuebuff string
		if err = rows.Scan(&entitybuff, &dialogbuff, &scriptbuff, &keybuff, &valuebuff); err != nil {
			return nil, errors.Wrap(err, "dialogs_scripts_data.scan")
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
		log.Add("Universe", universeID).Add("Dialogs", len(compiledDialogs)).Info("dialogs select")
	}

	return compiledDialogs, nil
}
