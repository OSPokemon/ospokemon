({
	class: 'menu',
	repeat: {},
	menus: [
		'menu/menubar/menubar',
		'menu/bindings/bindings',
		'menu/itembag/itembag',
		'menu/player',
		'menu/chat/chat',
		'menu/actions/actions',
		'menu/settings/settings',
		'menu/dialog/dialog',
		'menu/toaster/toaster'
	],
	build: function() {
		ospokemon.menu = this

		window.onbeforeunload = function (e) {
			return 'block'
		}

		$('body').keydown(this.keydown)
		$('body').keyup(this.keyup)

		$.each(this.menus, function(i, name) {
			ospokemon.element.build(name).then(function(menu) {
				$(ospokemon.menu).append(menu)
			})
		})
	},
	keydown: function(e) {
		if (!ospokemon.menu.repeat[e.key]) {
			ospokemon.menu.repeat[e.key] = true
			ospokemon.websocket.Send('Key.Down', e.key)
		}
	},
	keyup: function(e) {
		ospokemon.menu.repeat[e.key] = false
		ospokemon.websocket.Send('Key.Up', e.key)
	}
})