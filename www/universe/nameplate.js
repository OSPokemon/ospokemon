({
	class: 'universe.nameplate',
	build: function(entity) {
		this.entity = entity
		$(entity).append(this)
		return this
	},
	update: function(data) {
		if (this.username != data.username) {
			this.username = data.username
			$('.username', this).text(this.username)
		}

		var left = (this.entity.dx/2) - ($(this).width() / 2)
		if (left != this.left) {
			$(this).css({
				"left": left + "px"
			})
		}
	}
})