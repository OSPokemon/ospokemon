({
	class: 'menu.bindings.button',
	build: function(data) {
		$(this).droppable({
			'accept': '.menu-actions-button, .menu-bag-button',
			drop: this.drop
		})

		this.update(data)

		return this
	},
	update: function(data) {
		if ($(this).attr('onclick') != data.onclick) {
			$(this).attr('onclick', data.onclick)
		}

		var img = $('img', this)

		if (img.attr("src") != data.image) {
			img.attr("src", data.image)
		}

		var key = $('.key', this)

		if (key.text() != data.key) {
			this.key = data.key
			key.text(data.key)
		}

		var amount = $('.amount', this)

		if (amount.text() != data.amount) {
			amount.text(data.amount)
		}
	},
	drop: function(event, ui) {
		if (!event) {
			return
		}

		var bind = ui.draggable[0]

		if (bind.class == 'menu.bag.button') {
			ospokemon.websocket.Send('Binding.Set', {
				'key': this.key,
				'bagslot': bind.pos
			})
		} else if (bind.class == 'menu.actions.button') {
			console.log('wassup')
		} else {
			console.error('binding type not recognized')
		}
	}
})