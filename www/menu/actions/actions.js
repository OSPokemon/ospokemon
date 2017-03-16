({
	class: 'menu.actions',
	buttons: {},
	build: function() {
		ospokemon.menu.actions = this

		$(this).draggable().resizable()

		ospokemon.event.On('Update', this.update)
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
		$(ospokemon.menu.actions).slideDown('slow')

		ospokemon.element.get('menu/actions/button').then(function(el) {
			$.each(data.actions, function(i, action) {
				if (ospokemon.menu.actions.buttons[i]) {
					ospokemon.menu.actions.buttons[i].update(action)
				} else {
					ospokemon.menu.actions.buttons[i] = el.build(action)
					$('.panel-body', ospokemon.menu.actions).append(ospokemon.menu.actions.buttons[i])
				}
			})
		})
	}
})
