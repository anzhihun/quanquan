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
            var curDate = new Date();
            messageContainer.append(Mustache.render(MessageItemTemplate, {
                headImg: msg.HeadImg,
                name: msg.From,
                datetime: curDate.toLocaleDateString() + ' ' + curDate.toLocaleTimeString(),
                content: 'online'
            }));
//            UserView.addUser({
//                HeadImg: msg.HeadImg,
//                Name: msg.From
//            });
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
			var userList = global.mainframe.getUserListView().getUsers();
			var iconUrl = '';
			for(var index = 0, len = userList.length; index < len; index++) {
				if (userList.at(index).get('name') === msg.sender) {
					iconUrl = userList.at(index).get('iconUrl');
					break;
				}
			}
			
            var textMessage = new TalkMessage({
				user: {
					name: msg.sender,
					iconUrl: iconUrl
				},
				content: msg.content,
				dataTime: new Date().getTime()
			});
            
            messageBoard.getModel().add(textMessage);
        }
    }];
    
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