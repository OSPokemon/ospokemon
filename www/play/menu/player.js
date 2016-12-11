({
	class: 'menu.player',
	build: function() {
		ospokemon.menu.player = this

		$(this).draggable()

		ospokemon.event.On('Update', this.update)

		return this
	},
	toggle: function() {
		ospokemon.websocket.Send('Menu.Toggle', 'player')
	},
	update: function(data) {
		if (!data.player) {
			if (!$(ospokemon.menu.player).is(':hidden')) {
				$(ospokemon.menu.player).slideUp('slow')
			}
			return
		}

		if ($(ospokemon.menu.player).is(':hidden')) {
			$(ospokemon.menu.player).slideDown('slow')
		}

		$('.menu-player-name', ospokemon.menu.player).text(data.player.username)
		$('.menu-player-level', ospokemon.menu.player).text('Lv. '+data.player.level)
		$('.menu-player-money', ospokemon.menu.player).text('$'+data.player.money)
		$('.progress-bar', ospokemon.menu.player).attr("aria-valuenow", data.player.experience)
		$('.menu-player-experience span', ospokemon.menu.player).text(data.player.experience+'XP')
		// $('.menu-player-portrait', ospokemon.menu.player).attr("src", data.player.animations['portrait'])
	}
})