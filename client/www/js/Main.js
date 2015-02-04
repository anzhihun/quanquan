require.config({
    baseUrl: '/',
    paths: {
        text: 'js/thirdparty/text'
    }
});

var global = {
	
	/**
	 * @type {name:, iconUrl:, serverId: , online: true}
	 */
	currentUser: null,
	wsconn: null,
	/**
	 * @type {name: , isChannel: false}
	 */
	currentTalkTarget: null,
	
	/**
	 *@type {Mainframe}
	 */
	mainframe: null,
	
	/**
	 * @type {UserList}
	 * 
	 */
	allUsers: null,
    
    /**
     * @type {Object}
     * view strings
     */
    strings: ""
};

define(function (require) {
    'use strict';
    
    var WSConnector = require('js/WSConnector'),
        WSMsgHandler = require('js/WSMsgHandler'),

        LoginView = require('js/LoginView'),
		Mainframe = require('js/Mainframe');
	
    global.strings = require('/res/language');
    global.mainframe = new Mainframe();
    var loginView = new LoginView(function(){
        global.wsconn = createWebsocket(function(){
          global.mainframe.show();
            $(document).foundation();
        });
    }); 
	
    function createWebsocket(callbackOnConn) {
        var url = document.domain,
            wsUrl = 'ws://' + url + ':52013/rtmsg?id=' + global.currentUser.name ;
        return new WSConnector(wsUrl, new WSMsgHandler(), callbackOnConn);
    }
});