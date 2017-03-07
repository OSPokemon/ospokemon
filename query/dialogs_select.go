package query

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/game"
)

func DialogsSelect(entity *game.Entity, universe *game.Universe) (*game.Dialog, error) {
	rows, err := Connection.Query(
		"SELECT id, parent, lead, text, script FROM dialogs WHERE entity=? AND universe=?",
		entity.Id,
		universe.Id,
	)
	if err != nil {
		return nil, err
	}

	dialogs := make(map[uint]*game.Dialog)
	for rows.Next() {
		dialog := game.MakeDialog()
		if err = rows.Scan(&dialog.Id, &dialog.Parent, &dialog.Lead, &dialog.Text, &dialog.Script); err != nil {
			return nil, err
		}
		dialogs[dialog.Id] = dialog
	}
	rows.Close()

	rows, err = Connection.Query(
		"SELECT dialog, key, value FROM dialogs_data WHERE entity=? AND universe=?",
		entity.Id,
		universe.Id,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var dialogbuff uint
		var keybuff, valuebuff string
		if err = rows.Scan(&dialogbuff, &keybuff, &valuebuff); err != nil {
			return nil, err
		}
		dialogs[dialogbuff].Data[keybuff] = valuebuff
	}

	for _, dialog := range dialogs {
		if dialog.Id != 0 {
			parentDialog := dialogs[dialog.Parent]
			parentDialog.Choices = append(parentDialog.Choices, dialog)
		}
	}

	if len(dialogs) > 0 && err == nil {
		logrus.WithFields(logrus.Fields{
			"Universe": universe.Id,
			"Entity":   entity.Id,
			"Dialogs":  len(dialogs),
		}).Debug("dialogs select")
	}

	return dialogs[0], nil
}
