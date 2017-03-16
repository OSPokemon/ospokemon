({
	class: 'menu.dialog',
	data: {},
	build: function() {
		ospokemon.menu.dialog = this
		$(this).draggable()
		ospokemon.event.On('Update', this.update)
	},
	update: function(data) {
		if (!data.dialog) {
			if (!$(ospokemon.menu.dialog).is(':hidden')) {
				$(ospokemon.menu.dialog).slideUp('slow')
			}
			return
		}
		if ($(ospokemon.menu.dialog).is(':hidden')) {
			$(ospokemon.menu.dialog).slideDown('slow')
		}

		if (ospokemon.menu.dialog.log) {
			console.log(data)
			ospokemon.menu.dialog.log = false
		}

		var dialogdata = ospokemon.menu.dialog.data

		if (dialogdata.lead != data.dialog.lead) {
			dialogdata.lead = data.dialog.lead
			$("#menu-dialog-lead").html(ospokemon.username+": "+dialogdata.lead)
		}
		
		if (dialogdata.text != data.dialog.text) {
			dialogdata.text = data.dialog.text
			$("#menu-dialog-text").html(dialogdata.text)
		}

		var choices = data.dialog.choices.join("")
		if (dialogdata.choices != choices) {
			dialogdata.choices = choices

			ospokemon.element.get('menu/dialog/choice').then(function(el) {
				$('#menu-dialog-choices').html('')

				$.each(data.dialog.choices, function(i, choice) {
					choice = el.build(choice)
					$('#menu-dialog-choices').append(choice)
				})

				if (data.dialog.choices.length < 1) {
					$('#menu-dialog-choices').append(el.build('OK'))
				}
			})
		}
	}
})