({
	class: 'menu.bag.button',
	build: function(data) {
		$(this).draggable({ revert: "invalid" })

		this.update(data)

		return this
	},
	update: function(data) {
		if (data == null) {
			$(this).addClass('empty')
			return
		}
		$(this).removeClass('empty')

		if (this.pos != data.pos) {
			this.pos = data.pos
		}

		var span = $('span', this)
		if (span.text() != data.amount) {
			span.text(data.amount)
		}

		var img = $('img', this)
		if (img.attr("src") != data.item.image) {
			img.attr("src", data.item.image)
		}

		return this
	}
})
