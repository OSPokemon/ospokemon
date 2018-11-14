ospokemon.template = {
	cache: {},
	get: function(name) {
		var el = ospokemon.template.cache[name];
		if (!el) {
			return new Promise(function(resolve, reject) {
				ospokemon.template.load(name).then(resolve);
			})
		} else if (typeof el == 'Promise') {
			return el;
		};
		return new Promise(function(resolve, reject) {
			resolve(el);
		});
	},
	load: function(name) {
		ospokemon.template.cache[name] = new Promise(function(resolve, reject) {
			if (ospokemon.log.script) console.log('template request: ' + name + '.html');
			$.get(name + '.html')
			.done(function(data) {
				if (ospokemon.log.template) console.log('template loaded: ' + name);
				ospokemon.template.cache[name] = data;
			})
			.always(function() {
				ospokemon.template.cache[name] = ospokemon.template.cache[name] || 'error';
				resolve(ospokemon.template.cache[name]);
			});
		});
		return ospokemon.template.cache[name];
	}
};
