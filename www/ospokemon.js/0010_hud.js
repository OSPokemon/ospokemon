ospokemon.event.on('rootready', function() {
	ospokemon.hud = $('#hud')[0];
});

ospokemon.event.on('Update', function(data) {
	if (!ospokemon.hud) return;

	if (data.player) {
		$(ospokemon.hud).show();
		var entityid = data.entityid;
		var entity = data.universe[entityid];
		var hudPortrait = $('#hud-portrait');
		var hudName = $('#hud-name');

		if (hudPortrait.attr('src') != entity.image) {
			hudPortrait.attr('src', entity.image);
		};
		if (hudName.text() != entity.player.username) {
			hudName.text(entity.player.username)
		};
	} else {
		$(ospokemon.hud).hide();
	}
});
