define(function(require, exports, module){
	'use strict';
	
	var MessageInputView = Backbone.View.extend({
		el: '.send_area',
		initialize: function(){
			this.inputText = $('#msgInput')[0];
		},
		
		events: {
			'click #sendMsgBtn': 'sendTalk'
		},
		
		sendTalk: function(){
			var msg = {
				msgType: 'talk',
				contentType: 'text',
				sender: global.currentUser.name,
				is2P: !global.currentTalkTarget.isChannel,
				receiver: global.currentTalkTarget.name,
				content: this.inputText.value
			};
			global.wsconn.sendMessage(JSON.stringify(msg));
			this.inputText.value = '';
		}
	});
	
	return MessageInputView;
	
});