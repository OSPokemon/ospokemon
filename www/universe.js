({
	class: 'universe',
	entities: {},
	build: function(data) {
		this.data = data
		ospokemon.universe = this
		ospokemon.player = {}
		ospokemon.camera = {}
		ospokemon.event.On('Update', this.update)

		$('body').mousedown(this.mousedown)
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
			var entityDelete = {}
			$.each(ospokemon.universe.entities, function(key, val) {
				entityDelete[key] = true
			})

			$.each(ospokemon.universe.data, function(key, val) {
				entityDelete[key] = false

				if (!ospokemon.universe.entities[key]) {
					ospokemon.universe.entities[key] = el.build()
					$(ospokemon.universe).append(ospokemon.universe.entities[key])
				}

				ospokemon.universe.entities[key].update(val)
			})

			$.each(entityDelete, function(key, val) {
				if (val) {
					$(ospokemon.universe.entities[key]).remove()
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

		var minimumLeft = ($(window).width()/2) - 100
		var entityLeft = $(cameraEntity).offset().left
		if (entityLeft < minimumLeft) {
			$('.universe').css({
				left: (minimumLeft - cameraEntity.x) + 'px'
			})
		}

		var minimumTop = ($(window).height()/2) - 100
		var entityTop = $(cameraEntity).offset().top
		if (entityTop < minimumTop) {
			$('.universe').css({
				top: (minimumTop - cameraEntity.y) + 'px'
			})
		}

		var maximumLeft = ($(window).width()/2) + 100 - cameraEntity.dx
		if (entityLeft > maximumLeft) {
			$('.universe').css({
				left: (maximumLeft - cameraEntity.x ) + 'px'
			})
		}

		var maximumTop = ($(window).height()/2) + 100 - cameraEntity.dy
		if (entityTop > maximumTop) {
			$('.universe').css({
				top: (maximumTop - cameraEntity.y) + 'px'
			})
		}
	},
	mousedown: function(e) {
		if (e.button == 2) {
			var universeOffset = $('.universe').offset()
			ospokemon.websocket.Send('Click.Universe', {
				x: e.pageX - universeOffset.left,
				y: e.pageY - universeOffset.top
			})
		}
	}
})