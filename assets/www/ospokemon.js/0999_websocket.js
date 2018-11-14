ospokemon.websocket = {
	websocket: null,
	connect: function() {
		ospokemon.websocket.websocket = new WebSocket(window.location.protocol.replace('http', 'ws')+window.location.host+'/api/websocket');

		ospokemon.websocket.websocket.onmessage = function (e) {
			console.debug('websocket.onmessage', e);
			var data = JSON.parse(e.data);

			if (data.event == 'Update') {
				if (ospokemon.log.update) {
					console.log(data.data);
					ospokemon.log.update = false;
				};
			};

			ospokemon.event.fire(data.event, data.data);
		};

		ospokemon.websocket.websocket.onopen = function(e) {
			window.onbeforeunload = function (e) {
				return 'block';
			};
		};

		ospokemon.websocket.websocket.onclose = function(e) {
			window.onbeforeunload = null;
			if (ospokemon.username) {
				window.location.href = '/login/#' + ospokemon.username;
			} else {
				window.location.href = '/login/';
			}
		};
	},
	sendEvent: function(event, data) {
		if (!ospokemon.websocket.websocket) return;

		var msg;
		if (typeof data == "string") {
			msg = data;
		} else if (typeof data == "object") {
			msg = JSON.stringify(data);
		};

		if (ospokemon.log.websocket) console.log('send message: ' + event);

		ospokemon.websocket.websocket.send(JSON.stringify({
			"Event": event,
			"Message": msg
		}));
	}
};
