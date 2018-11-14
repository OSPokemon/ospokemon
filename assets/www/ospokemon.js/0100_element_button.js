var emptyData = {
	image: '',
	key: '',
	amount: ''
};

if (ospokemon.log.script) console.log('coded script: element/button');
ospokemon.script.cache['element/button'] = {
	build: function(data) {
		$(this).mousedown(this.mousedown);
		$(this).draggable({scroll: false, revert: "invalid"});

		if (data) {
			this.update(data)
		};
	},
	mousedown: function(event) {
		if (event.button == 2) {
			this.fire();
			event.stopPropagation();
			return false
		};
	},
	fire: function() {
		if (this.itemid) {
			ospokemon.websocket.sendEvent('ItemCast', this.itemid);
		} else if (this.spellid) {
			ospokemon.websocket.sendEvent('SpellCast', this.spellid);
		} else if (this.menuid) {
			ospokemon.websocket.sendEvent('MenuToggle', this.menuid);
		}
	},
	update: function(data) {
		if (!data) {
			data = emptyData;
		};

		if (this.image != data.image) {
			this.image = data.image;
			$('img', this).attr("src", this.image);
		};

		if (data.item && this.itemid != data.item.id) {
			this.itemid = data.item.id;
		} else if (data.spell && this.spellid != data.spell.id) {
			this.spellid = data.spell.id;
		} else if (data.menu && this.menuid != data.menu) {
			this.menuid = data.menu;
		};

		if (this.key != data.key) {
			this.key = data.key;
			$('.key', this).text(this.key.substr(0, 3));
		};

		if (this.amount != data.amount) {
			this.amount = data.amount;
			$('.amount', this).text(this.amount);
		};

		if (this.log) {
			console.log(data);
			this.log = false;
		};
	}
};
