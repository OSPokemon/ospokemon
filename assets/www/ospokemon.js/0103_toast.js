if (ospokemon.log.script) console.log('coded script: element/toast');
ospokemon.script.cache['element/toast'] = {
	build: function(data) {
		$(this).css({
			'color': data.color,
		});
		$('img', this).attr('src', data.image);
		$('span.message', this).html(data.message);
	}
};
