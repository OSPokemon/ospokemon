ospokemon.event.on('rootready', function() {
	ospokemon.toaster = $('#toaster')[0];

	ospokemon.event.on('Update', function(data) {
		var toaster = $(ospokemon.toaster);

		ospokemon.element.factory('element/toast').then(function(toastFactory) {
			$.each(data.toaster, function(toastid, toast) {
				var toast = toastFactory(toast);
				toaster.append(toast);

				$(toast).slideDown('slow');
				setTimeout(function() {
					$(toast).slideUp('slow');
				}, 3000);
			});
		});
	});
});
