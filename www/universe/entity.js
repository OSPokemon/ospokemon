({
	class: 'universe.entity',
	build: function() {
		$(this).mousedown(this.mousedown)

		return this
	},
	update: function(data) {
		var image = $('.entity-image', this)

		if (this.log) {
			console.log(data)
			this.log = false
		}

		this.id = data.id

		if (image.attr('src') != data.imaging.image) {
			image.attr('src', data.imaging.image)
		}

		if (this.x != data.shape.anchor.x || this.y != data.shape.anchor.y) {
			this.x = data.shape.anchor.x
			this.y = data.shape.anchor.y

			$(this).css({
				left: data.shape.anchor.x,
				top: data.shape.anchor.y
			})
		}

		if (this.dx != data.shape.dimension.dx || this.dy != data.shape.dimension.dy) {
			this.dx = data.shape.dimension.dx
			this.dy = data.shape.dimension.dy

			image.css({
				width: data.shape.dimension.dx,
				height: data.shape.dimension.dy
			})
		}

		if (data.player) {
			var entity = this
			ospokemon.element.get('universe/nameplate').then(function(el) {
				if (!entity.player) {
					entity.player = el.build(entity)
				}

				entity.player.update(data.player)
			})
		}
	},
	mousedown: function(e) {
		if (e.button == 2) {
			ospokemon.websocket.Send('Click.Entity', {
				entity: parseInt(this.id)
			})
		}
		return false
	}
})