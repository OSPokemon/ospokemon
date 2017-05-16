ospokemon.event.on('rootready', function() {
	ospokemon.menus = $('#menus')[0];
	ospokemon.menus.buttons = {};

	ospokemon.element.factory('element/button').then(function(buttonFactory) {
		$.each(['player', 'actions', 'itembag', 'settings'], function(i, menu) {
			ospokemon.menus.buttons[menu] = buttonFactory();
			ospokemon.menus.append(ospokemon.menus.buttons[menu]);
		});
	});
});

ospokemon.event.on('Update', function(data) {
	if (!ospokemon.menus) return;
	if (ospokemon.menus.log) {
		console.log(data);
		ospokemon.menus.log = false
	};

	$.each(ospokemon.menus.buttons, function(menu, button) {
		$.each(data.bindings, function(key, binding) {
			if (binding.menu == menu) {
				button.update(binding);
			};
		});
	});
});
