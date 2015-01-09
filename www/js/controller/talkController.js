/* global define, $, Mustache */
define(function(require, exports, module){
    'use strict';
    
    var MessageItemTemplate = require('text!/view/messageItem.html');
    var messageContainer = $('.main .message_area .body'),
		TextMessage = require('js/msg/TextMessage'),
		TextMessageView = require('js/msg/TextMessageView');
    
    // msg handlers 
    var handlers = [{
        msgType: 'join',
        handle: function(msg) {
			var textMessage = new TextMessage({
				user: {
					name: msg.From,
					iconUrl: msg.HeadImg
				},
				content: 'join',
				dataTime: new Date().getTime()
			});
			
			var textMessageView = new TextMessageView({
				model: textMessage
			});
			textMessageView.render();
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
            var curDate = new Date();
            messageContainer.append(Mustache.render(MessageItemTemplate, {
                headImg: msg.HeadImg,
                name: msg.From,
                datetime: curDate.toLocaleDateString() + ' ' + curDate.toLocaleTimeString(),
                content: msg.Content
            }));
        }
    }];
    
    function handle(msg) {
        
        if (messageContainer.length === 0) {
            messageContainer = $('.main .message_area .body');
        }
        
        for (var index = 0, len = handlers.length; index < len; index++) {
            if (handlers[index].msgType === msg.MsgType) {
                handlers[index].handle(msg);
                return;
            }
        }
    }
    
    function accept(msg) {
        if (msg === null || msg === undefined ) {
            return false;
        }
        
        for (var index = 0, len = handlers.length; index < len; index++) {
            if (handlers[index].msgType === msg.MsgType) {
                return true;
            }
        }
        
        return false;
    }
    
    exports.handle = handle;
    exports.accept = accept;
});