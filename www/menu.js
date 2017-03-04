({
	class: 'menu',
	repeat: {},
	build: function() {
		ospokemon.menu = this

		$('body').keydown(this.keydown)
		$('body').keyup(this.keyup)

		ospokemon.element.build('menu/menubar/menubar').then(function(menubar) {
			$(ospokemon.menu).append(menubar)
		})

		ospokemon.element.build('menu/bindings/bindings').then(function(bindings) {
			$(ospokemon.menu).append(bindings)
		})

		ospokemon.element.build('menu/itembag/itembag').then(function(itembag) {
			$(ospokemon.menu).append(itembag)
		})

		ospokemon.element.build('menu/player').then(function(player) {
			$(ospokemon.menu).append(player)
		})

		ospokemon.element.build('menu/actions/actions').then(function(actions) {
			$(ospokemon.menu).append(actions)
		})

		ospokemon.element.build('menu/settings/settings').then(function(settings) {
			$(ospokemon.menu).append(settings)
		})

		ospokemon.element.build('menu/dialog').then(function(dialog) {
			$(ospokemon.menu).append(dialog)
		})

		return this
	},
	keydown: function(e) {
		var key = e.key

		if (ospokemon.menu.repeat[key]) {
			return
		}

		ospokemon.menu.repeat[key] = true
		ospokemon.websocket.Send('Key.Down', key)
	},
	keyup: function(e) {
		var key = e.key
		ospokemon.menu.repeat[key] = false
		ospokemon.websocket.Send('Key.Up', key)
	}
})