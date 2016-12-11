({
	class: 'menu',
	repeat: {},
	build: function() {
		ospokemon.menu = this

		$('body').keydown(this.keydown)
		$('body').keyup(this.keyup)
		// ospokemon.event.On('Update', this.refresh)

		ospokemon.element.build('menu/menubar/menubar').then(function(menubar) {
			$(ospokemon.menu).append(menubar)
		})

		ospokemon.element.build('menu/bindings/bindings').then(function(bindings) {
			$(ospokemon.menu).append(bindings)
		})

		ospokemon.element.build('menu/bag/bag').then(function(bag) {
			$(ospokemon.menu).append(bag)
		})

		ospokemon.element.build('menu/player').then(function(player) {
			$(ospokemon.menu).append(player)
		})

		ospokemon.element.build('menu/actions/actions').then(function(actions) {
			$(ospokemon.menu).append(actions)
		})

		// ospokemon.BuildElement('menu/bindings/bindings').then(function(el) {
		// 	ospokemon.menu.bindings = el
		// 	$('body').append(el)
		// })
		// ospokemon.BuildElement('menu/actions/actions').then(function(el) {
		// 	ospokemon.menu.actions = el
		// 	$('body').append(el)
		// })
		// ospokemon.BuildElement('menu/player/player').then(function(el) {
		// 	ospokemon.menu.player = el
		// 	$('body').append(el)
		// })
		// ospokemon.BuildElement('menu/bag/bag').then(function(el) {
		// 	ospokemon.menu.bag = el
		// 	$('body').append(el)
		// })

		// ospokemon.BuildElement('menu/camera').then(function(el) {
		// 	ospokemon.camera = el
		// 	$('body').append(el)

		// })

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