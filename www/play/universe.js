({
	class: 'universe',
	entities: {},
	build: function(data) {
		this.data = data
		ospokemon.universe = this
		ospokemon.player = {}
		ospokemon.camera = {}
		ospokemon.event.On('Update', this.update)

		return this
	},
	update: function(data) {
		ospokemon.universe.data = data.universe

		if (ospokemon.universe.log) {
			console.log(ospokemon.universe.data)
			ospokemon.universe.log = false
		}

		if (ospokemon.camera.focus != data.entityid) {
			ospokemon.camera.focus = data.entityid
		}

		ospokemon.universe.focusCamera()

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
	},
	focusCamera: function() {
		var cameraEntity = ospokemon.universe.entities[ospokemon.camera.focus]

		if (!cameraEntity) {
			return
		}

		if (ospokemon.camera.log) {
			console.log(cameraEntity)
			ospokemon.camera.log = false
		}

		var minimumLeft = 200
		var entityLeft = $(cameraEntity).offset().left
		if (entityLeft < minimumLeft) {
			$('.universe').css({
				left: (minimumLeft - cameraEntity.data.shape.anchor.x) + 'px'
			})
		}

		var minimumTop = 200
		var entityTop = $(cameraEntity).offset().top
		if (entityTop < minimumTop) {
			$('.universe').css({
				top: (minimumTop - cameraEntity.data.shape.anchor.y) + 'px'
			})
		}

		var maximumLeft = $('.universe').width() - 200
		if (entityLeft > maximumLeft) {
			$('.universe').css({
				left: (maximumLeft - cameraEntity.data.shape.anchor.x ) + 'px'
			})
		}

		var maximumTop = $('.universe').height() - 250
		if (entityTop > maximumTop) {
			$('.universe').css({
				top: (maximumTop - cameraEntity.data.shape.anchor.y) + 'px'
			})
		}
	}
})