ospokemon.template.get('element/root').then(function(rootTpl) {
	ospokemon.root = $(rootTpl)[0];
	$('body').append(ospokemon.root);
	console.log('rootready');
	ospokemon.event.fire('rootready');
});
