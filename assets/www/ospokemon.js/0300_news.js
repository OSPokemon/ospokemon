ospokemon.event.on('rootready', function() {
	ospokemon.template.get('element/news').then(function(newsTpl) {
		ospokemon.news = $(newsTpl)[0];
		$('#root')[0].append(ospokemon.news);

		var connectCounter = 30;
		var connectAuto = true;
		var autoPlayCheckbox = $('#news-play-auto-checkbox')[0];

		function NewsCounter() {
			if (connectCounter == 0) {
				$(ospokemon.news).remove();
				ospokemon.websocket.connect();
			} else {
				connectCounter--;
				$('#news-play-auto-timer').html('Connecting... ' + connectCounter);
				if (autoPlayCheckbox.checked) window.setTimeout(NewsCounter, 1000);
			};
		};

		$('#news-play-button').click(function() {
			$(ospokemon.news).remove();
			ospokemon.websocket.connect();
		});

		$('#news-signup-button').click(function() {
			window.location.href = '/signup/';
		});

		$('#news-login-button').click(function() {
			window.location.href = '/login/';
		});

		$(autoPlayCheckbox).click(function(e) {
			if (autoPlayCheckbox.checked) {
				NewsCounter();
			};
		});

		$(autoPlayCheckbox).click();
	});
});
