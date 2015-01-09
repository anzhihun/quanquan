require.config({
    baseUrl: '/',
    paths: {
        text: 'js/thirdparty/text'
    }
});


define(function (require) {
    'use strict';

    var WSConnector = require('js/wsConnector'),
        WSMsgHandler = require('js/WSMsgHandler'),
        
        LoginView = require('js/LoginView'),
		Mainframe = require('js/Mainframe'),
        Context = require('js/context');
	
    var mainframe = new Mainframe();
    var loginView = new LoginView(function(){
        Context.wsconn = createWebsocket(function(){
          mainframe.show();
        });
    }); 
	
    function createWebsocket(callbackOnConn) {
        var url = document.domain,
            wsUrl = 'ws://' + url + ':52013/rtmsg?id=' + Context.currentUser ;
        return new WSConnector(wsUrl, new WSMsgHandler(), callbackOnConn);
    }
});