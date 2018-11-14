ospokemon.camera = {
	entityid: ''
};

ospokemon.event.on('Update', function(data) {
	if (!ospokemon.universe) return;

	if (data.entityid != ospokemon.camera.entityid) {
		ospokemon.camera.entityid = data.entityid
	};

	var cameraEntity = ospokemon.universe.entities[ospokemon.camera.entityid];
	if (!cameraEntity) return;

	if (ospokemon.camera.log) {
		console.log(cameraEntity);
		ospokemon.camera.log = false;
	};

	var minimumLeft = ($(window).width()/2) - 100;
	var entityLeft = $(cameraEntity).offset().left;
	if (entityLeft < minimumLeft) {
		$(ospokemon.universe).css({
			left: (minimumLeft - cameraEntity.x) + 'px'
		});
	};

	var minimumTop = ($(window).height()/2) - 100;
	var entityTop = $(cameraEntity).offset().top;
	if (entityTop < minimumTop) {
		$(ospokemon.universe).css({
			top: (minimumTop - cameraEntity.y) + 'px'
		});
	};

	var maximumLeft = ($(window).width()/2) + 100 - cameraEntity.dx;
	if (entityLeft > maximumLeft) {
		$(ospokemon.universe).css({
			left: (maximumLeft - cameraEntity.x ) + 'px'
		});
	};

	var maximumTop = ($(window).height()/2) + 100 - cameraEntity.dy;
	if (entityTop > maximumTop) {
		$(ospokemon.universe).css({
			top: (maximumTop - cameraEntity.y) + 'px'
		});
	};
});
