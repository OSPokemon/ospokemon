({
	class: 'menu.chat.history-item',
	build: function(data) {
		$('.username', this).html(data.username+':')
		$('.message', this).html(data.message)
		return this
	}
})