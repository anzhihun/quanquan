require.config({
    baseUrl: '/',
    paths: {
        text: 'js/thirdparty/text'
    }
});


define(function (require) {
    'use strict';

    var MainWindowHtml = require('text!/view/mainWindow.html'),
        LoginHtml = require('text!/view/login.html'),
//        TalkController = require('js/controller/talkController'),
        LoginController = require('js/controller/loginController'),
        WSConnector = require('js/wsConnector'),
        WSMsgHandler = require('js/WSMsgHandler'),
        ActionHandler = require('js/actionHandler'),
        
        UserRequester = require('js/requester/userRequester'),
        UserView = require('js/view/userView'),
        Context = require('js/context');

    document.body.innerHTML = LoginHtml;
    
    
    LoginController.bindActionHandler(switchToMainView);
    $(document).foundation();

    function createWebsocket() {
        var url = document.domain,
            wsUrl = 'ws://' + url + ':53240/rtmsg?id=' + Context.currentUser ;

        return new WSConnector(wsUrl, new WSMsgHandler());
    }

    function switchToMainView() {
        document.body.innerHTML = MainWindowHtml;
        $(document).foundation();
        
        var wsConnector = createWebsocket();
        ActionHandler.bindingHandler(wsConnector);
        UserRequester.getAllUser(UserView.updateAllUser);
    }

});