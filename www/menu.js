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

		$('body').keydown(this.keydown)
		$('body').keyup(this.keyup)

		$.each(this.menus, function(i, name) {
			ospokemon.element.build(name).then(function(menu) {
				$(ospokemon.menu).append(menu)
			})
		})
	},
	keydown: function(e) {
		var key = e.key
		if (!key) {
			key = String.fromCharCode(e.charCode).toLowerCase()
		}
		if (!key) {
			console.error("keydown detected, keycode unavailable. event follows.")
			console.error(e)
		}

		if (!ospokemon.menu.repeat[key]) {
			ospokemon.menu.repeat[key] = true
			ospokemon.websocket.Send('Key.Down', key)
		}
	},
	keyup: function(e) {
		ospokemon.menu.repeat[e.key] = false
		ospokemon.websocket.Send('Key.Up', e.key)
	}
})