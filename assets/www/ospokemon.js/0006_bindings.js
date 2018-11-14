ospokemon.event.on('rootready', function() {
	ospokemon.bindings = $('#bindings')[0];
	ospokemon.element.apply(ospokemon.bindings, {
		buttons: {},
		movement: {up:'',right:'',down:'',left:''},
		menus: {player:'',itembag:'',actions:'',settings:''},
		chat: ''
	});

	ospokemon.element.factory('element/button').then(function(buttonFactory) {
		for (var key = 1; key < 6; key++) {
			var button = ospokemon.bindings.buttons[key] = buttonFactory({key:key+'',image:'',amount:''});
			$(button).droppable({
				accept: '.menu-button',
				drop: function(event, ui) {
					if (!event) return;
					var bind = ui.draggable[0];
					if (bind) bind.style = "";

					if (bind.itemid) {
						ospokemon.websocket.sendEvent('BindingSet', {
							'key': this.key,
							'item': bind.itemid
						});
					} else if (bind.spell) {
						ospokemon.websocket.sendEvent('BindingSet', {
							'key': this.key,
							'spell': bind.spellid
						});
					} else {
						console.error('binding type not recognized');
					}
				}
			});
			ospokemon.bindings.append(button);
		}
	});
});

ospokemon.event.on('Update', function(data) {
	if (ospokemon.bindings.log) {
		console.log(data.bindings);
		ospokemon.bindings.log = false;
	};

	$.each(data.bindings, function(key, binding) {
		if (binding.walk) {
			ospokemon.bindings.movement[binding.walk] = binding.key;
		};

		if (binding.menu == 'chat' && key != ospokemon.bindings.chat) {
			ospokemon.bindings.chat = key;
			key = '';
		} else if (binding.menu) {
			ospokemon.bindings.menus[binding.menu] = binding.key;
		};

		if (ospokemon.bindings.buttons[key]) {
			ospokemon.bindings.buttons[key].update(binding);
		};
	});
});

ospokemon.event.on('keydown', function(key) {
	if (key == ospokemon.bindings.chat) {
		$('input', ospokemon.chat).focus();
	};
	$.each(ospokemon.bindings.movement, function(walk, walkKey) {
		if (key == walkKey) {
			ospokemon.websocket.sendEvent('MovementOn', walk);
		};
	});
	$.each(ospokemon.bindings.menus, function(menu, menuKey) {
		if (key == menuKey) {
			ospokemon.websocket.sendEvent('MenuToggle', menu);
		};
	});
});

ospokemon.event.on('keyup', function(key) {
	$.each(ospokemon.bindings.movement, function(walk, boundKey) {
		if (key == boundKey) {
			ospokemon.websocket.sendEvent('MovementOff', walk);
		};
	});
});
