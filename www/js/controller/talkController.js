/* global define, $, Mustache */
define(function(require, exports, module){
    'use strict';
    
    var MessageItemTemplate = require('text!/view/messageItem.html');
    var messageContainer = $('.main .message_area .body'),
		TalkMessage = require('js/msg/TalkMessage'),
		TalkMessageView = require('js/msg/TalkMessageView');
    
	var messageBoard = null;
    // msg handlers 
    var handlers = [{
        msgType: 'join',
        handle: function(msg) {
			var textMessage = new TalkMessage({
				user: {
					name: msg.From,
					iconUrl: msg.HeadImg
				},
				content: 'join',
				dataTime: new Date().getTime()
			});
            
            messageBoard.getModel().add(textMessage);
        }
    }, {
        msgType: 'online',
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
            
            messageBoard.getModel().add(textMessage);
        }
    }, {
        msgType: 'userLogin',
        handle: function(msg) {
            var msgContent = JSON.parse(msg.Content);
            var textMessage = new TalkMessage({
				user: {
					name: msgContent.name,
					iconUrl: msgContent.HeadImg
				},
				content: 'userLogin',
				dataTime: new Date().getTime()
			});
            
            messageBoard.getModel().add(textMessage);
        }
    }, {
        msgType: 'offline',
        handle: function(msg) {
            var curDate = new Date();
            messageContainer.append(Mustache.render(MessageItemTemplate, {
                headImg: msg.HeadImg,
                name: msg.From,
                datetime: curDate.toLocaleDateString() + ' ' + curDate.toLocaleTimeString(),
                content: 'offline'
            }));
        }
    }, {
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
            
            messageBoard.getModel().add(textMessage);
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
        
        if (messageContainer.length === 0) {
            messageContainer = $('.main .message_area .body');
        }
		
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