ospokemon.event.on('rootready', function() {
	ospokemon.settings = $('#settings')[0];
	$(ospokemon.settings).draggable();
	$('button', ospokemon.settings).click(function() {
		window.onbeforeunload = null;
		window.location.href = 'api/logout';
	});
	ospokemon.settings.toggle = function() {
		ospokemon.websocket.sendEvent('MenuToggle', 'settings');
	};

	var movementInputs = $('#settings-movement input', ospokemon.settings);
	$.each(['left', 'up', 'right', 'down'], function(i, b) {
		movementInputs[i].walk = b;
		$(movementInputs[i]).keydown(function(e) {
			ospokemon.websocket.sendEvent('BindingSet', {
				'key': event.key,
				'walk': this.walk
			});
			e.preventDefault();
			return false;
		});
	});

	var menuInputs = $('#settings-menus input', ospokemon.settings);
	$.each(['chat', 'player', 'itembag', 'actions', 'settings'], function(i, b) {
		menuInputs[i].menu = b;
		$(menuInputs[i]).keydown(function(e) {
			ospokemon.websocket.sendEvent('BindingSet', {
				'key': event.key,
				'menu': this.menu
			});
			e.preventDefault();
			return false;
		})
	});
});

ospokemon.event.on('Update', function(data) {
	if (!ospokemon.settings) return;

	if (data.settings) {
		$(ospokemon.settings).show();
	} else {
		$(ospokemon.settings).hide();
	};

	var menuInputs = $('#settings-menus input');
	var movementInputs = $('#settings-movement input');

	$.each(menuInputs, function(i, input) {
		$.each(ospokemon.bindings.menus, function(menu, key) {
			if (menu == input.menu && key != input.key) {
				input.key = key;
				$(input).val(key);
				return false;
			}
		});
	});

	$.each(movementInputs, function(i, input) {
		$.each(ospokemon.bindings.movement, function(walk, key) {
			if (walk == input.walk && key != input.key) {
				input.key = key;
				$(input).val(key);
				return false;
			}
		});
	});

	var chatInput = $('#settings-menus input')[0];
	if (chatInput.key != ospokemon.bindings.chat) {
		chatInput.key = ospokemon.bindings.chat;
		$(chatInput).val(chatInput.key);
	}
});