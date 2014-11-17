
define(function(require, exports){
    'use strict';
    
    var wsConnector = null;
    var Message = require('js/message'),
        ChannelRequester = require('js/requester/channelRequester');
    
    function bindingHandler(wsconn) {
        wsConnector = wsconn;
        $('#sendMsgBtn').click(sendMessage);
        $('#showAddChanDlgBtn').click(showAddChannelDlg);
        $('#addChannelButton').click(addNewChannel);
    }
    
    function sendMessage(){
        var msg = new Message();
        msg.Content = $('#msgInput')[0].value;
        wsConnector.sendMessage(JSON.stringify(msg));
        $('#msgInput')[0].value = '';
    }
    
    function showAddChannelDlg() {
        $('#addChannelDialog').foundation('reveal', 'open');
    }
    
    function addNewChannel() {
        var name = $('#newChannelNameInput')[0].value;
        ChannelRequester.addNewChannel(name, function(data, status){
            // add new channel to list
            
        });
        $('#addChannelDialog').foundation('reveal', 'close');
    }
    
    exports.bindingHandler = bindingHandler;
});