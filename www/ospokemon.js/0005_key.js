ospokemon.key = {
	cache: {},
	decode: function(e) {
		var key = e.key;
		if (!key) key = String.fromCharCode(e.charCode).toLowerCase();
		if (!key) console.error('unsupported charCode: ' + e.charCode);
		return key;
	}
};

$('body').keydown(function(e) {
	var key = ospokemon.key.decode(e);
	if (!ospokemon.key.cache[key]) {
		ospokemon.key.cache[key] = true;
		// ospokemon.websocket.Send('Key.Down', key)
		ospokemon.event.fire('keydown', key);
	};
});

$('body').keyup(function(e) {
	var key = ospokemon.key.decode(e);
	ospokemon.key.cache[key] = false;
	// ospokemon.websocket.Send('Key.Up', e.key);
	ospokemon.event.fire('keyup', key);
});
