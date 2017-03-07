package game

import (
	"github.com/ospokemon/ospokemon/part"
	// "time"
)

type Binding struct {
	Key string
	part.Parts
}

func MakeBinding() *Binding {
	return &Binding{
		Parts: make(part.Parts),
	}
}

func (binding *Binding) Part() string {
	return part.Binding
}

func (binding *Binding) SetAction(action *Action) {
	action.AddPart(binding)
	binding.Parts = action.Parts
}

func (binding *Binding) SetItemslot(itemslot *Itemslot) {
	itemslot.AddPart(binding)
	binding.Parts = itemslot.Parts
}

func (binding *Binding) SetWalk(walk string) {
	binding.Parts = make(part.Parts)
	binding.AddPart(Walk(walk))

	imaging := MakeImaging()
	imaging.Image = "/img/ui/walk/" + walk + ".png"
	binding.AddPart(imaging)
}

func (binding *Binding) SetMenu(menu string) {
	binding.Parts = make(part.Parts)
	binding.AddPart(Menu(menu))

	imaging := MakeImaging()
	imaging.Image = "/img/ui/menu/" + menu + ".png"
	binding.AddPart(imaging)
}

func (binding *Binding) RemoveParts() {
	if binding.Parts != nil {
		binding.RemovePart(binding)
		binding.Parts = nil
	}
}

// func (b *Binding) Update(u *Universe, e *Entity, d time.Duration) {
// }
