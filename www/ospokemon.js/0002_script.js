ospokemon.script = {
	cache: {},
	get: function(name) {
		var el = ospokemon.script.cache[name];
		if (!el) {
			return new Promise(function(resolve, reject) {
				ospokemon.script.load(name).then(resolve);
			});
		} else if (typeof el == 'Promise') {
			return el;
		};
		return new Promise(function(resolve, reject) {
			resolve(el);
		});
	},
	load: function(name) {
		ospokemon.script.cache[name] = new Promise(function(resolve, reject) {
			if (ospokemon.log.script) console.log('script request: ' + name + '.js');
			$.get(name + '.js')
			.done(function(data) {
				if (ospokemon.log.script) console.log('script loaded: ' + name);
				ospokemon.script.cache[name] = eval(data);
			})
			.always(function() {
				ospokemon.script.cache[name] = ospokemon.script.cache[name] || 'error';
				resolve(ospokemon.script.cache[name]);
			});
		});
		return ospokemon.script.cache[name];
	}
};
