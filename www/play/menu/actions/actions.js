({
	class: 'menu.actions',
	buttons: {},
	build: function() {
		ospokemon.menu.actions = this

		$(this).draggable().resizable()

		ospokemon.event.On('Update', this.update)

		return this
	},
	toggle: function() {
		ospokemon.websocket.Send('Menu.Toggle', 'actions')
	},
	update: function(data) {
		if (!data.actions) {
			if (!$(ospokemon.menu.actions).is(':hidden')) {
				$(ospokemon.menu.actions).slideUp('slow')
			}
			return
		}

		if ($(ospokemon.menu.actions).is(':hidden')) {
			$(ospokemon.menu.actions).slideDown('slow')
		}

		var actions = $('.panel-body', ospokemon.menu.actions)
		actions.empty()

		ospokemon.element.get('menu/actions/button').then(function(el) {
			$.each(data.actions, function(i, a) {
				var button = el.build(a)
				button.refresh()
				actions.append(button)
			})
		})

		return this
	}
})
