require.config({
    baseUrl: '/',
    paths: {
        text: 'js/thirdparty/text'
    }
});


define(function (require) {
    'use strict';

    var LoginHtml = require('text!/view/login.html'),
//        TalkController = require('js/controller/talkController'),
        LoginController = require('js/controller/loginController'),
        WSConnector = require('js/wsConnector'),
        WSMsgHandler = require('js/WSMsgHandler'),
        ActionHandler = require('js/actionHandler'),
        
        UserRequester = require('js/requester/userRequester'),
		Mainframe = require('js/Mainframe'),
		
        Context = require('js/context');
	
	var mainframe = new Mainframe();

    document.body.innerHTML = LoginHtml;

    LoginController.bindActionHandler(switchToMainView);
    $(document).foundation();

    function createWebsocket() {
        var url = document.domain,
            wsUrl = 'ws://' + url + ':52013/rtmsg?id=' + Context.currentUser ;

        return new WSConnector(wsUrl, new WSMsgHandler());
    }

    function switchToMainView() {
        var wsConnector = createWebsocket();
        ActionHandler.bindingHandler(wsConnector);
		
		mainframe.show();
    }

});