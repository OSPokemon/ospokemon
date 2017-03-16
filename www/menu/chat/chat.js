({
	class: 'menu.chat',
	repeat: {},
	build: function() {
		ospokemon.menu.chat = this

		ospokemon.event.On('Update', this.update)
		$('input', this)
			.keyup(this.keyup)
			.keydown(this.keydowninput)
			.focusin(this.focusin)
			.focusout(this.focusout)

		$(ospokemon.menu.chat).mousedown(this.mousedown)

		$('body').keydown(this.keydownglobal)
	},
	focusin: function(event) {
		$(ospokemon.menu.chat).addClass('active')
	},
	focusout: function(event) {
		$(ospokemon.menu.chat).removeClass('active')
	},
	keydownglobal: function(event) {
		var chat = ospokemon.menu.chat
		var input = $('input', chat)

		if (event.key == chat.key) {
			input.focus()
		}
	},
	keydowninput: function(event) {
		var chat = ospokemon.menu.chat
		var input = $('input', chat)

		if (event.key == 'Enter') {
			if (input.val() != '') {
				ospokemon.websocket.Send('Chat', input.val())
			}
			input.val('')
			input.blur()
		}
		event.stopPropagation()
	},
	keyup: function(event) {
		event.stopPropagation()
	},
	mousedown: function(event) {
		if (event.button == 2) {
			event.stopPropagation()
			$(ospokemon.menu.chat).hide()

			var target = document.elementFromPoint(event.clientX, event.clientY)
			while (!target.mousedown) {
				target = target.parentElement
			}
			target.mousedown(event)

			$(ospokemon.menu.chat).show()
			return false
		} else {
			setTimeout(function() {
				$('input', ospokemon.menu.chat).focus()
			})
			event.stopPropagation()
			return false
		}
	},
	update: function(data) {
		var chat = ospokemon.menu.chat
		var chatHistory = $('.chat-history', chat)

		$.each(data.bindings, function(key, binding) {
			if (binding.menu == 'chat' && chat.key != key) {
				chat.key = key
			}
		})

		ospokemon.element.get('menu/chat/history-item').then(function(el) {
			$.each(data.universe, function(entityid, entity) {
				if (entity.chat != chat.repeat[entityid]) {
					chat.repeat[entityid] = entity.chat

					if (!entity.chat || !entity.player) {
						return
					}

					chatHistory.append(el.build({
						username: entity.player.username,
						message: entity.chat
					}))
					chatHistory.scrollTop(chatHistory.prop('scrollHeight'))
				}
			})
		})
	}
})