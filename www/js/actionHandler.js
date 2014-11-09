/*global define, $ */
define(function(require, exports, module){
    var wsConnector = null;
    var Message = require('js/message');
    
    function bindingHandler(wsconn) {
        wsConnector = wsconn;
        $('#sendMsgBtn').click(sendMessage);
    }
    
    function sendMessage(event){
        var msg = new Message();
        msg.Content = $('#msgInput')[0].value;
        wsConnector.sendMessage(JSON.stringify(msg));
    }
    
    exports.bindingHandler = bindingHandler;
});