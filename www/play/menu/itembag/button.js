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

		if (this.log) {
			console.log(data)
			this.log = false
		}

		if (this.itemid !== data.item.id) {
			this.itemid = data.item.id
		}

		var span = $('span', this)
		if (span.text() != data.amount) {
			span.text(data.amount)
		}

		var img = $('img', this)
		if (img.attr("src") != data.imaging.image) {
			img.attr("src", data.imaging.image)
		}

		return this
	}
})
