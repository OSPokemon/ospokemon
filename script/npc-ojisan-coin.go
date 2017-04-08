package script

import (
	"ospokemon.com"
)

func init() {
	ospokemon.Scripts["npc-ojisan-coin"] = NPCOjisanCoin
}

var npcOjisanCoins = make(map[string]uint)

func NPCOjisanCoin(e *ospokemon.Entity, data map[string]interface{}) error {
	err := ItemChange(e, map[string]interface{}{"item": 1, "amount": -1})
	if err != nil {
		return err
	}

	player := e.GetPlayer()

	npcOjisanCoins[player.Username]++

	if npcOjisanCoins[player.Username] > 4 {
		dialog := ospokemon.MakeDialog()
		dialog.Lead = "Here's 1 coin"
		dialog.Text = "You are truly generous...<br/><br/>Please take this! Try it on!"
		player.AddPart(dialog)

		ItemChange(e, map[string]interface{}{"item": 2, "amount": 1})
	}

	return nil
}
