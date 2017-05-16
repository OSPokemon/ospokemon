if (ospokemon.log.script) console.log('coded script: element/dialog-choice');
ospokemon.script.cache['element/dialog-choice'] = {
	build: function(text) {
		this.text = text;
		$(this).click(this.click);
		$('span', this).html(text);
	},
	click: function() {
		ospokemon.websocket.sendEvent('Dialog.Choice', this.text)
	}
};
