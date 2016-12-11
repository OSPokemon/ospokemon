({
	class: 'menu.actions.button',
	build: function(data) {
		this.data = data
		$(this).draggable({ revert: "invalid" })
		return this
	},
	refresh: function() {
		$('span', this).text(this.data.spellid)
		$('img', this).attr("src", this.data.image)
		return this
	}
})
