/* global define, $, Mustache */
define(function(require, exports, module){
    'use strict';

    var TalkMessage = require('js/msg/TalkMessage');
    
    // msg handlers 
    var handlers = [{
        msgType: 'talk',
        handle: function(msg) {
			msg.content = msg.content.replace(new RegExp('\n', 'gm'), '<br>');
            var textMessage = new TalkMessage({
				user: {
					name: msg.sender,
					iconUrl: getUrl(msg.sender)
				},
				content: msg.content,
				dataTime: new Date().getTime()
			});
			var msgBoardId = '';
			if (msg.is2P) {
				msgBoardId = 'p2p::' + msg.receiver;
			} else {
				msgBoardId = 'chan::' + msg.receiver;
			}
            global.mainframe.getMessageBoard(msgBoardId).getModel().add(textMessage);
        }
    }];
    
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