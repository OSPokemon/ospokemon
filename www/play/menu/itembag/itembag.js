({
	class: 'menu.itembag',
	buttons: [],
	build: function() {
		ospokemon.menu.itembag = this

		$(this).draggable().resizable()

		ospokemon.event.On('Update', this.update)

		return this
	},
	toggle: function() {
		ospokemon.websocket.Send('Menu.Toggle', 'itembag')
	},
	update: function(data) {
		if (!data.itembag) {
			if (!$(ospokemon.menu.itembag).is(':hidden')) {
				$(ospokemon.menu.itembag).slideUp('slow')
			}
			return
		}

		ospokemon.menu.itembag.data = data

		if ($(ospokemon.menu.itembag).is(':hidden')) {
			$(ospokemon.menu.itembag).slideDown('slow')
		}

		var bag = $('.panel-body', ospokemon.menu.itembag)

		ospokemon.element.get('menu/itembag/button').then(function(el) {
			$.each(data.itembag, function(i, bdata) {
				if (ospokemon.menu.itembag.buttons[i]) {
					ospokemon.menu.itembag.buttons[i].update(bdata)
				} else {
					ospokemon.menu.itembag.buttons[i] = el.build(bdata)
					bag.append(ospokemon.menu.itembag.buttons[i])
				}
			})
		})

		return this
	}
})
