ospokemon.event.on('rootready', function() {
	ospokemon.actions = $('#actions')[0];

	$(ospokemon.actions).draggable().resizable();

	ospokemon.actions.buttons = {};
	ospokemon.actions.toggle = function() {
		ospokemon.websocket.sendEvent('MenuToggle', 'actions');
	}
});

ospokemon.event.on('Update', function(data) {
	if (!ospokemon.actions) return;

	if (data.actions) {
		$(ospokemon.actions).show();
	} else {
		$(ospokemon.actions).hide();
	};

	ospokemon.element.factory('element/button').then(function(buttonFactory) {
		$.each(data.actions, function(i, action) {
			if (ospokemon.actions.buttons[i]) {
				ospokemon.actions.buttons[i].update(action);
			} else {
				ospokemon.actions.buttons[i] = buttonFactory(action);
				$('.panel-body', ospokemon.actions).append(ospokemon.actions.buttons[i]);
			}
		});
	});
});
