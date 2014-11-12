/*global define, $ */
define(function(require, exports, module){
    var wsConnector = null;
    var Message = require('js/message'),
        ChannelRequester = require('js/requester/channelRequester');
    
    function bindingHandler(wsconn) {
        wsConnector = wsconn;
        $('#sendMsgBtn').click(sendMessage);
        $('#showAddChanDlgBtn').click(showAddChannelDlg);
        $('#addChannelButton').click(addNewChannel);
    }
    
    function sendMessage(event){
        var msg = new Message();
        msg.Content = $('#msgInput')[0].value;
        wsConnector.sendMessage(JSON.stringify(msg));
        $('#msgInput')[0].value = '';
    }
    
    function showAddChannelDlg(event) {
        $('#addChannelDialog').foundation('reveal', 'open');
    }
    
    function addNewChannel(event) {
        var name = $('#newChannelNameInput')[0].value;
        ChannelRequester.addNewChannel(name, function(data, status){
            // add new channel to list
            
        });
        $('#addChannelDialog').foundation('reveal', 'close');
    }
    
    exports.bindingHandler = bindingHandler;
});