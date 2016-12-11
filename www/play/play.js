ospokemon = {}

ospokemon.element = {
	elements: {},
	get: function(name) {
		var element = ospokemon.element.elements[name]
		if (element && element.script != null && element.html != null) {
			return new Promise(function(resolve, reject) {
				resolve(element)
			})
		} else {
			ospokemon.element._init(name)
			return new Promise(function(resolve, reject) {
				ospokemon.element.loadHtml(name).then(function(el) {
					ospokemon.element.loadScript(name).then(function(el) {
						resolve(el)
					})
				})
			})
		}
	},
	build: function(name) {
		return new Promise(function(resolve, reject) {
			ospokemon.element.get(name).then(function(el) {
				resolve(el.build())
			})
		})
	},
	loadHtml: function(name) {
		ospokemon.element._init(name)

		return new Promise(function(resolve, reject) {
			$.get(name + '.html')
			.done(function(data) {
				ospokemon.element.elements[name].html = data
			})
			.always(function() {
				ospokemon.element.elements[name].html = ospokemon.element.elements[name].html || false
				resolve(ospokemon.element.elements[name])
			})
		})
	},
	loadScript: function(name) {
		ospokemon.element._init(name)

		return new Promise(function(resolve, reject) {
			$.get(name + '.js')
			.done(function(data) {
				ospokemon.element.elements[name].script = eval(data)
			})
			.always(function() {
				ospokemon.element.elements[name].script = ospokemon.element.elements[name].script || false
				resolve(ospokemon.element.elements[name])
			})
		})
	},
	_init: function(name) {
		if (ospokemon.element.elements[name]) {
			return
		}

		ospokemon.element.elements[name] = {
			class: name,
			html: null,
			script: null,
			build: function(data) {
				var el = ospokemon.element.elements[name]
				var html = $(el.html)[0]

				$.each(el.script, function(key, val) {
					html[key] = val
				})

				if (!html) {
					console.log(el)
					return
				}
				if (!html.build) {
					console.log(el)
					return
				}

				return html.build(data)
			}
		}
	}
}

ospokemon.event = {}

ospokemon.event.On = function(event, f) {
	ospokemon.event[event] = ospokemon.event[event] || []
	ospokemon.event[event].push(f)
	return ospokemon.event[event].length - 1
}

ospokemon.event.Off = function(event, id) {
	ospokemon.event[event][id] = false
}

ospokemon.event.Fire = function() {
	var args = Array.prototype.slice.call(arguments)
	var event = args.shift()

	if (!ospokemon.event[event]) {
		// console.log('no handlers for event: '+event)
		return
	}

	for (var i = 0; i < ospokemon.event[event].length; i++) {
		var f = ospokemon.event[event][i]

		if (f) {
			f.apply(null, args)
		}
	}
}

ospokemon.websocket = new WebSocket('ws://' + window.location.host + '/api/websocket')

ospokemon.websocket.onmessage = function (e) {
	var data = JSON.parse(e.data)

	if (data.event == 'Update' && data.data.username != ospokemon.username) {
		ospokemon.username = data.data.username
	}

	ospokemon.event.Fire(data.event, data.data)
}

ospokemon.websocket.onclose = function(e) {
	window.location.href = '/login'
}

ospokemon.websocket.Send = function(event, cmd) {
	var msg = ""

	if (typeof cmd == "string") {
		msg = cmd
	} else if (typeof cmd == "object") {
		msg = JSON.stringify(cmd)
	}

	ospokemon.websocket.send(JSON.stringify({
		"Username": ospokemon.username,
		"Event": event,
		"Message": msg
	}))
}

ospokemon.element.build('universe').then(function(el) {
	$('body').append(el)
})
ospokemon.element.build('menu').then(function(el) {
	$('body').append(el)
})

// ospokemon.BuildElement('menu/button').then(function(el) {
// 	$('body').append(el)
// })

// ospokemon.event.On("Universe.Update", function(data) {
// 	ospokemon.player.data = data.player
// })