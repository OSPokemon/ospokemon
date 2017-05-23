var chatItemFactory = false;

ospokemon.element.factory('element/chat-item').then(function(f) {
	chatItemFactory = f;
});

ospokemon.event.on('rootready', function() {
	ospokemon.chat = $('#chat')[0];
	ospokemon.chat.repeat = {};

	$(ospokemon.chat).mousedown(function(event) {
		if (event.button == 2) {
			return;
		} else {
			setTimeout(function() {
				$('input', ospokemon.chat).focus()
			});
			event.stopPropagation();
			return false
		}
	});

	var chatInput = $('input', ospokemon.chat);

	chatInput.keyup(function(event) {
		event.stopPropagation();
	});
	chatInput.focusin(function(event) {
		$(ospokemon.chat).addClass('active');
	});
	chatInput.focusout(function(event) {
		$(ospokemon.chat).removeClass('active');
	});
	chatInput.keydown(function(event) {
		if (event.key == 'Enter') {
			if (chatInput.val() != '') ospokemon.websocket.sendEvent('Chat', chatInput.val());
			chatInput.val('');
			chatInput.blur();
		};
		event.stopPropagation();
	});
});

ospokemon.event.on('Update', function(data) {
	if (!ospokemon.chat) return;
	if (!chatItemFactory) return;

	var chatHistory = $('.chat-history', ospokemon.chat);

	$.each(data.universe, function(entityid, entity) {
		if (entity.chat != chat.repeat[entityid]) {
			ospokemon.chat.repeat[entityid] = entity.chat;

			if (!entity.chat || !entity.player) return;

			chatHistory.append(chatItemFactory({
				username: entity.player.username,
				message: entity.chat
			}));
			chatHistory.scrollTop(chatHistory.prop('scrollHeight'));
		}
	});
});
