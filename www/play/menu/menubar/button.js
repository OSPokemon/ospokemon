({
	class: 'menu.menubar.button',
	build: function(data) {
		this.data = data

		if (this.data) {
			this.refresh()
		}

		return this
	},
	refresh: function() {
		if ($(this).attr('onclick') != this.data.onclick) {
			$(this).attr('onclick', this.data.onclick)
		}

		var img = $('img', this)

		if (img.attr("src") != this.data.img) {
			img.attr("src", this.data.img)
		}

		var span = $('span', this)

		if (span.text() != this.data.text) {
			span.text(this.data.text)
		}
	}
})