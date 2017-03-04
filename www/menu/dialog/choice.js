({
	class: 'menu.dialog.choice',
	build: function(text) {
		$(this).click(this.click)

		this.text = text
		$('span', this).html(text)

		return this
	},
	click: function() {
		ospokemon.websocket.Send('Dialog.Choice', this.text)
	}
})