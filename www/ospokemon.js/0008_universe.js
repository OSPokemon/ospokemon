ospokemon.universe = false;
var entityFactory = false;

ospokemon.element.factory('element/entity').then(function(f) {
	entityFactory = f;
});
ospokemon.event.on('rootready', function() {
	ospokemon.universe = $('#universe')[0];
	ospokemon.universe.entities = {};
});

ospokemon.event.on('Update', function(data) {
	if (!ospokemon.universe) return;
	if (!entityFactory) return;

	ospokemon.universe.data = data.universe;

	if (ospokemon.universe.log) {
		console.log(ospokemon.universe.data);
		ospokemon.universe.log = false;
	};

	var entityDelete = {};
	$.each(ospokemon.universe.entities, function(key, val) {
		entityDelete[key] = true;
	});

	$.each(ospokemon.universe.data, function(key, val) {
		entityDelete[key] = false;

		if (!ospokemon.universe.entities[key]) {
			ospokemon.universe.entities[key] = entityFactory();
			ospokemon.universe.append(ospokemon.universe.entities[key]);
		};

		ospokemon.universe.entities[key].update(val);
	});

	$.each(entityDelete, function(key, val) {
		if (val) $(ospokemon.universe.entities[key]).remove();
	});
});

$('body').mousedown(function(e) {
	if (e.button == 2) {
		var universeOffset = $(ospokemon.universe).offset();
		ospokemon.websocket.sendEvent('Click.Universe', {
			x: e.pageX - universeOffset.left,
			y: e.pageY - universeOffset.top
		});
	};
});
