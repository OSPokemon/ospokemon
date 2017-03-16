({
	class: 'menu.dialog.choice',
	build: function(text) {
		$(this).click(this.click)

		this.text = text
		$('span', this).html(text)
	},
	click: function() {
		ospokemon.websocket.Send('Dialog.Choice', this.text)
	}
})