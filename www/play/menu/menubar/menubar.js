({
	class: 'menu.menubar',
	build: function() {
		ospokemon.menu.menubar = this

		ospokemon.element.get('menu/menubar/button').then(function(el) {

			var playerbutton = el.build({
				text: 'c',
				img: '/img/ui/player.png',
				onclick: 'ospokemon.menu.player.toggle()'
			})
			$(ospokemon.menu.menubar).append(playerbutton)

			var bagbutton = el.build({
				text: 'b',
				img: '/img/ui/bag.png',
				onclick: 'ospokemon.menu.bag.toggle()'
			})
			$(ospokemon.menu.menubar).append(bagbutton)

			var mapbutton = el.build({
				text: 'm',
				img: '/img/ui/map.png',
				onclick: 'ospokemon.menu.map.toggle()'
			})
			$(ospokemon.menu.menubar).append(mapbutton)

			var movesbutton = el.build({
				text: 'v',
				img: '/img/ui/moves.png',
				onclick: 'ospokemon.menu.actions.toggle()'
			})
			$(ospokemon.menu.menubar).append(movesbutton)

			var pokeballsbutton = el.build({
				text: 'x',
				img: '/img/ui/pokeballs.png',
				onclick: 'ospokemon.menu.pokemon.toggle()'
			})
			$(ospokemon.menu.menubar).append(pokeballsbutton)

			var settingsbutton = el.build({
				text: 'z',
				img: '/img/ui/settings.png',
				onclick: 'ospokemon.menu.settings.toggle()'
			})
			$(ospokemon.menu.menubar).append(settingsbutton)

		})

		return this
	},
	refresh: function() {
	}
})