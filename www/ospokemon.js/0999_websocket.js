ospokemon.websocket = new WebSocket(window.location.protocol.replace('http', 'ws')+window.location.host+'/api/websocket');

ospokemon.websocket.onmessage = function (e) {
	var data = JSON.parse(e.data);

	if (data.event == 'Update') {
		if (ospokemon.log.update) {
			console.log(data.data);
			ospokemon.log.update = false;
		};
	};

	ospokemon.event.fire(data.event, data.data);
};

ospokemon.websocket.onopen = function(e) {
	window.onbeforeunload = function (e) {
		return 'block';
	};
};

ospokemon.websocket.onclose = function(e) {
	window.onbeforeunload = null;
	if (ospokemon.username) {
		window.location.href = '/login/#' + ospokemon.username;
	} else {
		window.location.href = '/login/';
	}
};

ospokemon.websocket.sendEvent = function(event, data) {
	var msg;
	if (typeof data == "string") {
		msg = data;
	} else if (typeof data == "object") {
		msg = JSON.stringify(data);
	};

	if (ospokemon.log.websocket) console.log('send message: ' + event);
	ospokemon.websocket.send(JSON.stringify({
		"Event": event,
		"Message": msg
	}));
};
