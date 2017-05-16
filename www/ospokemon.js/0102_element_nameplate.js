if (ospokemon.log.script) console.log('coded script: element/nameplate');
ospokemon.script.cache['element/nameplate'] = {
	build: function(entity) {
		this.entity = entity;
		entity.append(this);
	},
	update: function(data) {
		if (this.log) {
			console.log(data);
			this.log = false;
		};

		if (this.username != data.player.username) {
			this.username = data.player.username;
			$('.username', this).html(this.username);
		};

		if (this.chat != data.chat) {
			this.chat = data.chat;
			if (this.chat) {
				$('.username', this).html(this.username + ':');
				$('.message', this).html(this.chat);
			} else {
				$('.username', this).html(this.username);
				$('.message', this).html('');
			};
		};

		var left = (this.entity.dx/2) - ($(this).width() / 2);
		if (left != this.left) {
			$(this).css({
				'left': left + 'px'
			});
		};
	}
};