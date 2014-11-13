require.config({
    baseUrl: '/',
    paths: {
        text: 'js/thirdparty/text'
    }
});


define(function (require, exports, module) {
    'use strict';

    var MainWindowHtml = require('text!/view/mainWindow.html'),
        LoginHtml = require('text!/view/login.html'),
        TalkController = require('js/controller/talkController'),
        LoginController = require('js/controller/loginController'),
        WSConnector = require('js/wsConnector'),
        WSMsgHandler = require('js/WSMsgHandler'),
        ActionHandler = require('js/actionHandler'),
        
        UserRequester = require('js/requester/userRequester'),
        UserView = require('js/view/userView');

    document.body.innerHTML = LoginHtml;
    var wsConnector = createWebsocket();

    function createWebsocket() {
        var url = document.domain,
            wsUrl = 'ws://' + url + ':53240/rtmsg';

        return new WSConnector(wsUrl, new WSMsgHandler());
    }
    
    LoginController.bindActionHandler();
    $(document).foundation();
    
//    ActionHandler.bindingHandler(wsConnector);
//    UserRequester.getAllUser(UserView.updateAllUser);
    
    
//    TalkController.handle();
//    TalkController.handle();
});