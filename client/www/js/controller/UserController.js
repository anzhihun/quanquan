define(function (require, exports) {
	'use strict';

	var TalkMessage = require('js/msg/TalkMessage');
	
	function getUrl(userName) {
		var userList = global.allUsers;
		var iconUrl = '/images/defaultHead.png';
		for (var index = 0, len = userList.length; index < len; index++) {
			if (userList.at(index).get('name') === userName) {
				iconUrl = userList.at(index).get('iconUrl');
				break;
			}
		}
		return iconUrl;
	}

	// msg handlers 
	var handlers = [{
		msgType: 'newUser',
		handle: function (msg) {
			global.mainframe.getUserListView().getUsers().add(msg.user);
		}
    }, {
		msgType: 'online',
		handle: function (msg) {
			msg.content = msg.content.replace(new RegExp('\n', 'gm'), '<br>');
			var textMessage = new TalkMessage({
				user: {
					name: msg.sender,
					iconUrl: getUrl(msg.sender)
				},
				content: msg.content,
				dataTime: new Date().getTime()
			});
			
			// future: we should send msg to all channel which the user is belong to 
			global.mainframe.getMessageBoard('chan::Global').getModel().add(textMessage);
		}
    }, {
		msgType: 'join',
		handle: function (msg) {

		}
    }, {
		msgType: 'offline',
		handle: function (msg) {

		}
    }];

	function handle(msg) {

		for (var index = 0, len = handlers.length; index < len; index++) {
			if (handlers[index].msgType === msg.msgType) {
				handlers[index].handle(msg);
				return true;
			}
		}

		return false;
	}

	exports.handle = handle;
});