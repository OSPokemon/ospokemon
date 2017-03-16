package query

import (
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
)

func DialogsSelect(universe *game.Universe) (map[uint]*game.Dialog, error) {
	rows, err := Connection.Query(
		"SELECT entity, id, parent, lead, text, script FROM dialogs WHERE universe=?",
		universe.Id,
	)
	if err != nil {
		return nil, err
	}

	dialogs := make(map[uint]map[uint]*game.Dialog)
	for rows.Next() {
		var entityId uint
		dialog := game.MakeDialog()
		if err = rows.Scan(&entityId, &dialog.Id, &dialog.Parent, &dialog.Lead, &dialog.Text, &dialog.Script); err != nil {
			return nil, err
		}

		if dialogs[entityId] == nil {
			dialogs[entityId] = make(map[uint]*game.Dialog)
		}
		dialogs[entityId][dialog.Id] = dialog
	}
	rows.Close()

	rows, err = Connection.Query(
		"SELECT entity, dialog, key, value FROM dialogs_data WHERE universe=?",
		universe.Id,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var entityId, dialogbuff uint
		var keybuff, valuebuff string
		if err = rows.Scan(&entityId, &dialogbuff, &keybuff, &valuebuff); err != nil {
			return nil, err
		}
		dialogs[entityId][dialogbuff].Data[keybuff] = valuebuff
	}

	for _, entityDialogs := range dialogs {
		for _, dialog := range entityDialogs {
			if dialog.Id != 0 {
				parentDialog := entityDialogs[dialog.Parent]
				parentDialog.Choices = append(parentDialog.Choices, dialog)
			}
		}
	}

	compiledDialogs := make(map[uint]*game.Dialog)
	for entityId, entityDialogs := range dialogs {
		compiledDialogs[entityId] = entityDialogs[0]
	}

	if len(compiledDialogs) > 0 && err == nil {
		log.Add("Universe", universe.Id).Add("Dialogs", len(compiledDialogs)).Debug("dialogs select")
	}

	return compiledDialogs, nil
}
