if (ospokemon.log.script) console.log('coded script: element/chat-item');
ospokemon.script.cache['element/chat-item'] = {
	build: function(data) {
		$('.username', this).html(data.username+':');
		$('.message', this).html(data.message);
	}
};
