ospokemon.event.on('rootready', function() {
	ospokemon.itembag = $('#itembag')[0];
	ospokemon.itembag.buttons = [];
	$(ospokemon.itembag).draggable().resizable();

	ospokemon.itembag.toggle = function() {
		ospokemon.websocket.sendEvent('MenuToggle', 'itembag');
	}
});

ospokemon.event.on('Update', function(data) {
	if (!ospokemon.itembag) return;

	if (data.itembag) {
		$(ospokemon.itembag).show();
		var bag = $('.panel-body', ospokemon.itembag);

		ospokemon.element.factory('element/button').then(function(buttonFactory) {
			$.each(data.itembag, function(i, bdata) {
				if (ospokemon.itembag.buttons[i]) {
					ospokemon.itembag.buttons[i].update(bdata);
				} else {
					ospokemon.itembag.buttons[i] = buttonFactory(bdata);
					bag.append(ospokemon.itembag.buttons[i]);
				}
			});
		});
	} else {
		$(ospokemon.itembag).hide();
	}
});
