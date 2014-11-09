/* global define, $, Mustache */
define(function(require, exports, module){
    'use strict';
    
    var MessageItemTemplate = require('text!/view/messageItem.html');
    var messageContainer = $('.main .message_area .body');
    
    // msg handlers 
    var handlers = [{
        msgType: 'join',
        handle: function(msg) {
            messageContainer.append(Mustache.render(MessageItemTemplate, {
                headImg: '/images/anzhihun.png',
                name: 'anzhihun',
                datetime: '2014-08-12 12:00:34',
                content: 'message content'
            }));
        }
    }, {
        msgType: 'online',
        handle: function(msg) {
            messageContainer.append(Mustache.render(MessageItemTemplate, {
                headImg: '/images/anzhihun.png',
                name: 'anzhihun',
                datetime: '2014-08-12 12:00:34',
                content: 'message content'
            }));
        }
    }, {
        msgType: 'offline',
        handle: function(msg) {
            messageContainer.append(Mustache.render(MessageItemTemplate, {
                headImg: '/images/anzhihun.png',
                name: 'anzhihun',
                datetime: '2014-08-12 12:00:34',
                content: 'message content'
            }));
        }
    }, {
        msgType: 'talk',
        handle: function(msg) {
            messageContainer.append(Mustache.render(MessageItemTemplate, {
                headImg: '/images/anzhihun.png',
                name: 'anzhihun',
                datetime: '2014-08-12 12:00:34',
                content: 'message content'
            }));
        }
    }];
    
    function handle(msg) {
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