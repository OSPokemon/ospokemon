({
	class: 'menu.toaster.toast',
	build: function(data) {
		$(this).css({
			'color': data.color,
		})

		$('img', this).attr('src', data.image)

		$('span.message', this).html(data.message)

		return this
	}
})