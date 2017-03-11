({
	class: 'universe.nameplate',
	build: function(entity) {
		this.entity = entity
		$(entity).append(this)
		return this
	},
	update: function(data) {
		if (this.log) {
			console.log(data)
			this.log = false
		}

		if (this.username != data.player.username) {
			this.username = data.player.username
			$('.username', this).html(this.username)
		}

		if (this.message != data.chat) {
			this.message = data.chat
			if (this.message) {
				$('.username', this).html(this.username + ':')
				$('.message', this).html(this.message)
			} else {
				$('.username', this).html(this.username)
				$('.message', this).html('')

			}
		}

		var left = (this.entity.dx/2) - ($(this).width() / 2)
		if (left != this.left) {
			$(this).css({
				"left": left + "px"
			})
		}
	}
})