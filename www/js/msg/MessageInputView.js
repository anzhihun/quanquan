define(function(require, exports, module){
	'use strict';
	
	var TalkMessage = require('js/msg/TalkMessage');
	
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
			// send to server
			global.wsconn.sendMessage(JSON.stringify(msg));
			
			// add message to local dialogue board.
			if (!global.currentTalkTarget.isChannel) {
				var textMessage = new TalkMessage({
					user: global.currentUser,
					content: this.inputText.value,
					dataTime: new Date().getTime()
				});
				var msgBoardId =  'p2p::' + global.currentTalkTarget.name;
				global.mainframe.getMessageBoard(msgBoardId).getModel().add(textMessage);
			}
			
			this.inputText.value = '';
		},
		
		clear: function(){
			this.stopListening();
			this.off();
			this.undelegateEvents();
		}
	});
	
	return MessageInputView;
	
});