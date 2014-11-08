/* global define, $, Mustache */
define(function(require, exports, module){
    'use strict';
    
    var MessageItemTemplate = require('text!/view/messageItem.html');
    
    function receiveTalk(msg) {
        var messageContainer = $('.main .message_area .body');
        messageContainer.append(Mustache.render(MessageItemTemplate, {
            headImg: '/images/anzhihun.png',
            name: 'anzhihun',
            datetime: '2014-08-12 12:00:34',
            content: 'message content'
        }));
    }
    
    exports.receiveTalk = receiveTalk;
});