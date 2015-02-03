define(function (require, exports, module) {
	'use strict';
	var NewChannelDialog = Backbone.View.extend({
		initialize: function (container) {
			this.$el = $(container).find('#addChannelDialog');
			this.newChanInput = this.$el.find('#newChannelNameInput')[0];
			this.newChanInput.value = '';
		},

		events: {
			'click .button': 'addNewChannel'
		},

		open: function () {
			this.newChanInput.value = '';
			this.$el.foundation('reveal', 'open');
		},

		close: function () {
			this.$el.foundation('reveal', 'close');
		},

		addNewChannel: function () {
			var self = this;
			var channelName = this.newChanInput.value;
			$.post('/channel', JSON.stringify({
				name: channelName,
				creator: global.currentUser.name
			})).done(function (channel) {
				self.close();
			}).fail(function (obj) {
				//TODO  show errors
//				$('#signUpPanel .error')[0].innerHTML = obj.responseText;
//				$('#signUpPanel .error').show();
				self.close();
			});
		}
	});

	return NewChannelDialog;

});