({
	class: 'menu.menubar',
	keys: ['z', 'x', 'c', 'v', 'b'],
	buttons: {},
	build: function() {
		ospokemon.menu.menubar = this

		ospokemon.element.get('menu/button').then(function(el) {
			for (var keyi = 0; keyi < ospokemon.menu.menubar.keys.length; keyi++) {
				var key = ospokemon.menu.menubar.keys[keyi]

				var button = el.build({
					key: key,
					imaging: {}
				})
				$(ospokemon.menu.menubar).append(button)
				ospokemon.menu.menubar.buttons[key] = button
			}
		})

		ospokemon.event.On('Update', this.update)
	},
	update: function(data) {
		if (ospokemon.menu.menubar.log) {
			console.log(data)
			ospokemon.menu.menubar.log = false
		}

		if (ospokemon.menu.menubar.buttons.length < 1) {
			return
		}

		for (var keyi = 0; keyi < ospokemon.menu.menubar.keys.length; keyi++) {
			var key = ospokemon.menu.menubar.keys[keyi]

			if (data.bindings[key]) {
				if (!ospokemon.menu.menubar.buttons[key]) {
					console.log("no button found for keybinding: " + key)
					return
				}
				ospokemon.menu.menubar.buttons[key].update(data.bindings[key])
			}
		}
	}
})