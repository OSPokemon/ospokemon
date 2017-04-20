({
	class: 'menu.button',
	emptyData: {
		imaging: {
			image: ''
		},
		key: '',
		amount: ''
	},
	build: function(data) {
		$(this).mousedown(this.mousedown)
		$(this).draggable({scroll: false, revert: "invalid"})

		if (data) {
			this.update(data)
		}
	},
	update: function(data) {
		if (!data) {
			data = ospokemon.element.elements['menu/button'].script.emptyData
		}

		if (this.image != data.imaging.image) {
			this.image = data.imaging.image
			$('img', this).attr("src", this.image)
		}

		if (data.item && this.itemid != data.item.id) {
			this.itemid = data.item.id
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
	mousedown: function(event) {
		if (event.button == 2) {
			this.fire()
			event.stopPropagation()
			return false
		}
	},
	fire: function() {
		if (this.itemid) {
			ospokemon.websocket.Send('Item.Cast', {'item': this.itemid+''})
			return
		}
	}
})
