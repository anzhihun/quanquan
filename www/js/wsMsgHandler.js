define(function (require) {
    'use strict';

    function WSMsgHandler() {
    }

    // 接收实时消息
    WSMsgHandler.prototype.onMessage = function (msg) {
        
        console.log('receive server msg: ' + msg);
        if (msg === null || msg.trim().length === 0) {
            return;
        }
        
        var msgObj = JSON.parse(msg);
        switch (msgObj.MsgType) {
        case 'ProtocolData':
            console.log('receive protocol msg: ' + msg);
            // this._handlerProtocolMsg(msgObj);
            break;
        case 'hi':
            this._handleHelloMsg(msgObj);
            break;
        default:
            console.warn('未知消息' + msg);
        }
    };

    WSMsgHandler.prototype._handleHelloMsg = function (msg) {
        $('.main .message_area .body').append('<span>' + msg.From + '</span>: <span>' + msg.Content + '</span>');
    };

    // 处理协议消息

    return WSMsgHandler;
});