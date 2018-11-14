ospokemon.event.on('rootready', function() {
	ospokemon.dialog = $('#dialog')[0];
	$(ospokemon.dialog).draggable();
});

ospokemon.event.on('Update', function(data) {
	if (!ospokemon.dialog) return;

	if (data.dialog) {
		$(ospokemon.dialog).show();
	} else {
		$(ospokemon.dialog).hide();
		return;
	};

	if (ospokemon.dialog.log) {
		console.log(data);
		ospokemon.dialog.log = false;
	};

	if (ospokemon.dialog.lead != data.dialog.lead) {
		ospokemon.dialog.lead = data.dialog.lead;
		$("#dialog-lead").html(ospokemon.username+": "+ospokemon.dialog.lead);
	};

	if (dialogdata.text != data.dialog.text) {
		dialogdata.text = data.dialog.text;
		$("#dialog-text").html(dialogdata.text);
	};

	var choices = data.dialog.choices.join("");
	if (dialogdata.choices != choices) {
		dialogdata.choices = choices;

		ospokemon.element.factory('element/dialog-choice').then(function(choiceFactory) {
			$('#dialog-choices').html('');

			$.each(data.dialog.choices, function(i, choice) {
				choice = choiceFactory(choice);
				$('#dialog-choices').append(choice);
			});

			if (data.dialog.choices.length < 1) {
				$('#dialog-choices').append(choiceFactory('OK'));
			};
		});
	};
});
