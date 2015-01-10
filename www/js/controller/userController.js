define(function (require, exports) {
	'use strict';

	var TalkMessage = require('js/msg/TalkMessage');
	
	function getUrl(userName) {
		var userList = global.mainframe.getUserListView().getUsers();
		var iconUrl = '/images/defaultHead.png';
		for (var index = 0, len = userList.length; index < len; index++) {
			if (userList.at(index).get('name') === userName) {
				iconUrl = userList.at(index).get('iconUrl');
				break;
			}
		}
		return iconUrl;
	}

	var messageBoard = null;
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

			messageBoard.getModel().add(textMessage);
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

		messageBoard = messageBoard || global.mainframe.getMessageBoard();

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