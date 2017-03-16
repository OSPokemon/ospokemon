({
	class: 'menu.toaster',
	build: function() {
		ospokemon.menu.toaster = this
		ospokemon.event.On('Update', this.update)
	},
	update: function(data) {
		var toaster = $(ospokemon.menu.toaster)

		ospokemon.element.get('menu/toaster/toast').then(function(el) {
			$.each(data.toaster, function(toastid, toast) {
				var toast = el.build(toast)
				toaster.append(toast)

				$(toast).slideDown('slow')
				setTimeout(function() {
					$(toast).slideUp('slow')
				}, 3000)
			})
		})
	}
})