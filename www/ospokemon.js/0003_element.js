ospokemon.element = {
	cache: {},
	factory: function(name) {
		if (!ospokemon.element.cache[name]) {
			return new Promise(function(resolve, reject) {
				ospokemon.element.load(name).then(resolve);
			});
		} else if (typeof ospokemon.element.cache[name] == 'Promise') {
			return ospokemon.element.cache[name];
		};
		return new Promise(function(resolve, reject) {
			resolve(ospokemon.element.cache[name]);
		});
	},
	load: function(name) {
		ospokemon.element.cache[name] = new Promise(function(resolve, reject) {
			ospokemon.template.get(name).then(function(template) {
				ospokemon.script.get(name).then(function(script) {
					ospokemon.element.cache[name] = function() {
						var html = $(template)[0];
						ospokemon.element.apply(html, script);
						if (html && html.build) html.build.apply(html, arguments);
						return html;
					};
					resolve(ospokemon.element.cache[name]);
				});
			});
		});
		return ospokemon.element.cache[name];
	},
	build: function(name) {
		return new Promise(function(resolve, reject) {
			ospokemon.element.factory(name).then(function(factory) {
				resolve(factory());
			});
		});
	},
	apply: function(base, props) {
		$.each(props, function(k, v) {base[k] = v;});
	}
};
