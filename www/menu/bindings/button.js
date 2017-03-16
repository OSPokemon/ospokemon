({
	class: 'menu.bindings.button',
	build: function(data) {
		$(this).droppable({
			'accept': '.menu-actions-button, .menu-bag-button',
			drop: this.drop
		})

		this.update(data)
	},
	update: function(data) {
		if ($(this).attr('onclick') != data.onclick) {
			$(this).attr('onclick', data.onclick)
		}

		var img = $('img', this)

		if (this.log) {
			console.log(data)
			this.log = false
		}

		var key = $('.key', this)
		if (key.text() != data.key) {
			this.key = data.key
			key.text(data.key)
		}

		if (data.imaging) {
			if (img.attr("src") != data.imaging.image) {
				img.attr("src", data.imaging.image)
			}
		}

		if (data.itemslot) {
			var amount = $('.amount', this)

			if (amount.text() != data.itemslot.amount) {
				amount.text(data.itemslot.amount)
			}
		} else if (data.action) {
		}
	},
	drop: function(event, ui) {
		if (!event) {
			return
		}

		var bind = ui.draggable[0]

		if (bind.class == 'menu.bag.button') {
			bind.style = ""
			ospokemon.websocket.Send('Binding.Set', {
				'key': this.key,
				'item': bind.itemid
			})
		} else if (bind.class == 'menu.actions.button') {
			bind.style = ""
			ospokemon.websocket.Send('Binding.Set', {
				'key': this.key,
				'spell': bind.spell
			})
		} else {
			console.error('binding type not recognized')
		}
	}
})