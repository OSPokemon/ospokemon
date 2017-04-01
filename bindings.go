package ospokemon

import (
	"ospokemon.com/json"
	// "time"
)

const PARTbindings = "bindings"

type Bindings map[string]*Binding

func (b Bindings) Part() string {
	return PARTbindings
}

func (parts Parts) GetBindings() Bindings {
	bindings, _ := parts[PARTbindings].(Bindings)
	return bindings
}

func (bindings Bindings) SetAction(key string, action *Action) {
	var binding *Binding

	if binding = bindings[key]; binding != nil {
		binding.RemoveParts()
	} else {
		binding = MakeBinding()
		binding.Key = key
	}

	binding.SetAction(action)
	bindings[binding.Key] = binding
}

func (bindings Bindings) SetItemslot(key string, itemslot *Itemslot) {
	var binding *Binding

	if binding = bindings[key]; binding != nil {
		binding.RemoveParts()
	} else {
		binding = MakeBinding()
		binding.Key = key
	}

	binding.SetItemslot(itemslot)
	bindings[binding.Key] = binding
}

func (bindings Bindings) SetWalk(key string, walk string) {
	var binding *Binding

	if binding = bindings[key]; binding != nil {
		binding.RemoveParts()
	} else {
		binding = MakeBinding()
		binding.Key = key
	}

	binding.SetWalk(walk)
	bindings[binding.Key] = binding
}

func (bindings Bindings) SetMenu(key string, menu string) error {
	var binding *Binding

	if binding = bindings[key]; binding != nil {
		binding.RemoveParts()
	} else {
		binding = MakeBinding()
		binding.Key = key
	}

	binding.SetMenu(menu)
	bindings[binding.Key] = binding
	return nil
}

func (bindings Bindings) Remove(key string) error {
	if binding := bindings[key]; binding != nil {
		binding.RemoveParts()
	}

	delete(bindings, key)
	return nil
}

func (bindings Bindings) Json() json.Json {
	json := json.Json{}
	for key, binding := range bindings {
		json[key] = binding.Json()
	}
	return json
}

// func (b Bindings) Update(u *Universe, e *Entity, d time.Duration) {
// 	for _, binding := range b {
// 		binding.Update(u, e, d)
// 	}
// }
