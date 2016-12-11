({
	class: 'universe.entity',
	build: function(data) {
		this.data = data

		return this
	},
	refresh: function() {
		if (!this.data.comp) {
			console.log(this)
			return
		}

		var image = $('.entity-image', this)

		if (image.attr('src') != this.data.image) {
			image.attr('src', this.data.image)
		}

		image.css({
			width: this.data.comp.location.shape.dimension.dx,
			height: this.data.comp.location.shape.dimension.dy
		})
		$(this).css({
			left: this.data.comp.location.shape.anchor.x,
			top: this.data.comp.location.shape.anchor.y
		})
	}
})