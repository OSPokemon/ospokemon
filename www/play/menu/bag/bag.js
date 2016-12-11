({
	class: 'menu.bag',
	buttons: [],
	build: function() {
		ospokemon.menu.bag = this

		$(this).draggable().resizable()

		ospokemon.event.On('Update', this.update)

		return this
	},
	toggle: function() {
		ospokemon.websocket.Send('Menu.Toggle', 'bag')
	},
	update: function(data) {
		if (!data.bag) {
			if (!$(ospokemon.menu.bag).is(':hidden')) {
				$(ospokemon.menu.bag).slideUp('slow')
			}
			return
		}

		if ($(ospokemon.menu.bag).is(':hidden')) {
			$(ospokemon.menu.bag).slideDown('slow')
		}

		var bag = $('.panel-body', ospokemon.menu.bag)

		ospokemon.element.get('menu/bag/button').then(function(el) {
			$.each(data.bag, function(i, bdata) {
				if (ospokemon.menu.bag.buttons[i]) {
					ospokemon.menu.bag.buttons[i].update(bdata)
				} else {
					ospokemon.menu.bag.buttons[i] = el.build(bdata)
					bag.append(ospokemon.menu.bag.buttons[i])
				}
			})
		})

		return this
	}
})
