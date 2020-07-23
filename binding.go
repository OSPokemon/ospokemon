package ospokemon

import "taylz.io/types"

const PARTbinding = "binding"

type Binding struct {
	Key string
	Parts
}

func MakeBinding() *Binding {
	return &Binding{
		Parts: make(Parts),
	}
}

func (binding *Binding) Part() string {
	return PARTbinding
}

func (parts Parts) GetBinding() *Binding {
	binding, _ := parts[PARTbinding].(*Binding)
	return binding
}

func (binding *Binding) SetAction(action *Action) {
	action.AddPart(action)
	action.AddPart(binding)
	binding.Parts = action.Parts
}

func (binding *Binding) SetItemslot(itemslot *Itemslot) {
	itemslot.AddPart(itemslot)
	itemslot.AddPart(binding)
	binding.Parts = itemslot.Parts
}

func (binding *Binding) SetWalk(walk string) {
	binding.Parts = make(Parts)
	binding.AddPart(Walk(walk))

	imaging := MakeImaging()
	imaging.Image = "/img/ui/walk/" + walk + ".png"
	binding.AddPart(imaging)
}

func (binding *Binding) SetMenu(menu string) {
	binding.Parts = make(Parts)
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

func (binding *Binding) Json() types.Dict {
	json := types.Dict{
		"key": binding.Key,
	}

	if imaging := binding.GetImaging(); imaging != nil {
		json["image"] = imaging.Image
	}

	if walk := binding.GetWalk(); walk != "" {
		json["walk"] = walk
	} else if menu := binding.GetMenu(); menu != "" {
		json["menu"] = menu
	} else if itemslot := binding.GetItemslot(); itemslot != nil {
		json["item"] = itemslot.Json()
	}

	return json
}
