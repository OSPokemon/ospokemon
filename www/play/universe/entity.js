({
	class: 'universe.entity',
	build: function(data) {
		this.data = data

		return this
	},
	refresh: function() {
		if (!this.data) {
			console.log(this)
			return
		}

		var image = $('.entity-image', this)

		if (image.attr('src') != this.data.imaging.image) {
			image.attr('src', this.data.imaging.image)
		}

		image.css({
			width: this.data.shape.dimension.dx,
			height: this.data.shape.dimension.dy
		})
		$(this).css({
			left: this.data.shape.anchor.x,
			top: this.data.shape.anchor.y
		})
	}
})