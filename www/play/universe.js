({
	class: 'universe',
	entities: {},
	build: function(data) {
		this.data = data
		ospokemon.universe = this
		ospokemon.player = {}
		ospokemon.event.On('Update', this.update)
		return this
	},
	update: function(data) {
		ospokemon.universe.data = data.universe

		ospokemon.element.get('universe/entity').then(function(el) {
			$.each(ospokemon.universe.data, function(key, val) {
				if (!val) {
					if (ospokemon.universe.entities[key]) {
						$(ospokemon.universe.entities[key]).remove()
						ospokemon.universe.entities[key] = null
					}
				} else if (ospokemon.universe.entities[key]) {
					ospokemon.universe.entities[key].data = val
					ospokemon.universe.entities[key].refresh()
				} else {
					ospokemon.universe.entities[key] = el.build(val)
					$(ospokemon.universe).append(ospokemon.universe.entities[key])
				}
			})
		})
	}
})