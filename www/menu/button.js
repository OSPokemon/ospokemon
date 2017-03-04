({
	class: 'menu.button',
	emptyData: {
		imaging: {
			image: ''
		},
		key: '',
		amount: '',
	},
	build: function(data) {
		$(this).bind('click', this, this.click)
		$(this).draggable({ scroll: false, revert: "invalid" })

		if (data) {
			this.update(data)
		}

		return this
	},
	update: function(data) {
		if (!data) {
			data = ospokemon.element.elements['menu/button'].script.emptyData
		}

		if (this.image != data.imaging.image) {
			this.image = data.imaging.image
			$('img', this).attr("src", this.image)
		}

		if (this.key != data.key) {
			this.key = data.key
			$('.key', this).text(this.key)
		}

		if (this.amount != data.amount) {
			this.amount = data.amount

			$('.amount', this).text(this.amount)
		}

		if (this.log) {
			console.log(data)
			this.log = false
		}
	},
	click: function(button) {
	}
})