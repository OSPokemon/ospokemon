if (ospokemon.log.script) console.log('coded script: element/entity');
ospokemon.script.cache['element/entity'] = {
	build: function() {
		$(this).mousedown(this.mousedown);
	},
	update: function(data) {
		this.id = data.id;

		if (this.log) {
			console.log(data);
			this.log = false;
		};

		var image = $('.entity-image', this);
		if (image.attr('src') != data.image) {
			image.attr('src', data.image);
		};

		if (this.x != data.shape.anchor.x || this.y != data.shape.anchor.y) {
			this.x = data.shape.anchor.x;
			this.y = data.shape.anchor.y;

			$(this).css({
				left: data.shape.anchor.x,
				top: data.shape.anchor.y
			});
		};

		if (this.dx != data.shape.dimension.dx || this.dy != data.shape.dimension.dy) {
			this.dx = data.shape.dimension.dx;
			this.dy = data.shape.dimension.dy;

			image.css({
				width: this.dx,
				height: this.dy
			});
		};

		var entity = this;

		ospokemon.element.factory('element/nameplate').then(function(nameplateFactory) {
			if (data.player) {
				if (!entity.nameplate) {
					entity.nameplate = nameplateFactory(entity);
				};

				entity.nameplate.update(data);
			}
		});
	},
	mousedown: function(e) {
		if (e.button == 2) {
			ospokemon.websocket.sendEvent('Click.Entity', {
				entity: parseInt(this.id)
			});
		};
		return false;
	}
};