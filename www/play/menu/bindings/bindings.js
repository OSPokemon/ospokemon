({
	class: 'menu.bindings',
	buttons: {},
	build: function() {
		ospokemon.menu.bindings = this

		ospokemon.element.get('menu/bindings/button').then(function(el) {
			for (var bindingcount = 0; bindingcount < 10; bindingcount++) {
				var key = (bindingcount + 1) % 10 // 1, 2, 3, ..., 0

				var data = {
					key: key+'',
					spellid: false,
					itemid: false,
					timer: 0,
					image: false
				}

				ospokemon.menu.bindings.buttons[key] = el.build(data)
				$(ospokemon.menu.bindings).append(ospokemon.menu.bindings.buttons[key])
			}

			ospokemon.event.On('Update', this.update)
		})

		return this
	},
	update: function(data) {
		for (var bindingcount = 0; bindingcount < 10; bindingcount++) {
			var key = (bindingcount + 1) % 10 // 1, 2, 3, ..., 0
			if (data.bindings[key]) {
				ospokemon.menu.bindings.buttons[key].update(data.bindings[key])
			}
		}
	}
})
