({
	class: 'menu.actions.button',
	build: function(data) {
		$(this).draggable({ revert: "invalid" })
		this.update(data)
	},
	update: function(data) {
		if (this.log) {
			console.log(data)
			this.log = false
		}

		if (this.spell != data.spell.id) {
			this.spell = data.spell.id
		}

		return

		var img = $('img', this)
		if (img.attr("src") != data.imaging.image) {
			img.attr("src", data.imaging.image)
		}
	}
})
