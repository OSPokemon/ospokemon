package ospokemon

type DialogTest interface {
	Test(*Player) bool
}

type DialogItemTest struct {
	Item   uint
	Amount int
}

func (test *DialogItemTest) Test(player *Player) bool {
	itembag := player.GetItembag()
	return itembag.GetItems()[test.Item] >= test.Amount
}
