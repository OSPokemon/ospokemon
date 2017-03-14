({
	class: 'menu.settings',
	build: function() {
		ospokemon.menu.settings = this

		$(this).draggable()
		$('button', this).click(this.logout)

		var movementInputs = $('#menus-settings-movement input', this)
		$.each(['left', 'up', 'right', 'down'], function(i, b) {
			var direction = b
			$(movementInputs[i]).keydown(function(e) {
				ospokemon.websocket.Send('Binding.Set', {
					'key': event.key,
					'walk': direction
				})
				e.preventDefault()
				return false
			})
		})

		var menuInputs = $('#menus-settings-menus input', this)
		$.each(['chat', 'player', 'itembag', 'actions', 'settings'], function(i, b) {
			var menu = b
			$(menuInputs[i]).keydown(function(e) {
				ospokemon.websocket.Send('Binding.Set', {
					'key': event.key,
					'menu': menu
				})
				e.preventDefault()
				return false
			})
		})

		$('input', this).on('keydown', function() {
			return false
		})
		$('input', this).on('keyup', function() {
			return false
		})

		$('#menus-settings-logout .btn-danger', this).click(this.logout)

		ospokemon.event.On('Update', this.update)

		return this
	},
	toggle: function() {
		ospokemon.websocket.Send('Menu.Toggle', 'settings')
	},
	update: function(data) {
		if (!data.settings) {
			if (!$(ospokemon.menu.settings).is(':hidden')) {
				$(ospokemon.menu.settings).slideUp('slow')
			}
			return
		}

		if ($(ospokemon.menu.settings).is(':hidden')) {
			$(ospokemon.menu.settings).slideDown('slow')
		}

		var menuInputs = $('#menus-settings-menus input')
		var movementInputs = $('#menus-settings-movement input')

		$.each(data.bindings, function(key, val) {
			// console.log(this)
			if (val.walk == 'left' && ospokemon.menu.settings.walkleft != key) {
				ospokemon.menu.settings.walkleft = key
				$(movementInputs[0]).val(key)
			} else if (val.walk == 'up' && ospokemon.menu.settings.walkup != key) {
				ospokemon.menu.settings.walkup = key
				$(movementInputs[1]).val(key)
			} else if (val.walk == 'right' && ospokemon.menu.settings.walkright != key) {
				ospokemon.menu.settings.walkright = key
				$(movementInputs[2]).val(key)
			} else if (val.walk == 'down' && ospokemon.menu.settings.walkdown != key) {
				ospokemon.menu.settings.walkdown = key
				$(movementInputs[3]).val(key)
			} else if (val.menu == 'chat' && ospokemon.menu.settings.menuchat != key) {
				ospokemon.menu.settings.menuchat = key
				$(menuInputs[0]).val(key)
			}else if (val.menu == 'player' && ospokemon.menu.settings.menuplayer != key) {
				ospokemon.menu.settings.menuplayer = key
				$(menuInputs[1]).val(key)
			}else if (val.menu == 'itembag' && ospokemon.menu.settings.menuitembag != key) {
				ospokemon.menu.settings.menuitembag = key
				$(menuInputs[2]).val(key)
			}else if (val.menu == 'actions' && ospokemon.menu.settings.menuactions != key) {
				ospokemon.menu.settings.menuactions = key
				$(menuInputs[3]).val(key)
			}else if (val.menu == 'settings' && ospokemon.menu.settings.menusettings != key) {
				ospokemon.menu.settings.menusettings = key
				$(menuInputs[4]).val(key)
			}
		})
	},
	keydown: function(e) {

	},
	logout: function() {
		window.onbeforeunload = null
		window.location.href = 'api/logout'
	}
})