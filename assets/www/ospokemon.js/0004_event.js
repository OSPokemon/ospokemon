ospokemon.event = {
	cache: {},
	on: function(event, f) {
		ospokemon.event.cache[event] = ospokemon.event.cache[event] || [];
		ospokemon.event.cache[event].push(f);
		return ospokemon.event.cache[event].length - 1;
	},
	off: function(event, id) {
		ospokemon.event.cache[event][id] = false;
	},
	fire: function() {
		var args = Array.prototype.slice.call(arguments);
		var event = args.shift();
		var eventList = ospokemon.event.cache[event];
		if (!eventList) return;
		for (var i = 0; i < eventList.length; i++) {
			var f = eventList[i];

			if (f) {
				f.apply(null, args);
			}
		}
	}
};
