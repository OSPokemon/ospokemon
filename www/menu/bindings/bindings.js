({
	class: 'menu.bindings',
	buttons: {},
	build: function() {
		ospokemon.menu.bindings = this

		ospokemon.element.get('menu/button').then(function(el) {
			for (var bindingcount = 0; bindingcount < 10; bindingcount++) {
				var key = (bindingcount + 1) % 10 // 1, 2, 3, ..., 0

				var data = {
					key: key+'',
					imaging: {image: ''},
					amount: ''
				}

				var button = ospokemon.menu.bindings.buttons[key] = el.build(data)
				$(button).droppable({accept: '.menu-button', drop: ospokemon.menu.bindings.drop})
				$(ospokemon.menu.bindings).append(button)
			}
		})

		ospokemon.event.On('Update', this.update)
	},
	update: function(data) {
		if (ospokemon.menu.bindings.log) {
			console.log(data)
			ospokemon.menu.bindings.log = false
		}

		if (!ospokemon.menu.bindings.buttons['0']) {
			return
		}

		for (var bindingcount = 0; bindingcount < 10; bindingcount++) {
			var key = (bindingcount + 1) % 10 // 1, 2, 3, ..., 0
			if (data.bindings[key]) {
				ospokemon.menu.bindings.buttons[key].update(data.bindings[key])
			}
		}
	},
	drop: function(event, ui) {
		if (!event) {
			return
		}

		var bind = ui.draggable[0]

		if (bind.itemid) {
			bind.style = ""

			ospokemon.websocket.Send('Binding.Set', {
				'key': this.key,
				'item': bind.itemid
			})
		} else if (bind.spell) {
			bind.style = ""

			ospokemon.websocket.Send('Binding.Set', {
				'key': this.key,
				'spell': bind.spell
			})
		} else {
			console.error('binding type not recognized')
		}
	}
})
