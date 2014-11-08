/* global define, require, document */

require.config({
    baseUrl: '/',
    paths: {
        text: 'js/thirdparty/text'
    }
});


define(function (require, exports, module) {
    'use strict';

    var MainWindowHtml = require('text!/view/mainWindow.html'),
        TalkController = require('js/controller/talkController'),
        WSConnector = require('js/wsConnector'),
        WSMsgHandler = require('js/WSMsgHandler');

    document.body.innerHTML = MainWindowHtml;
    var wsConnector = createWebsocket();

    function createWebsocket() {
        var url = document.domain,
            wsUrl = 'ws://' + url + ':53240/rtmsg';

        return new WSConnector(wsUrl, new WSMsgHandler());
    }
    
    TalkController.receiveTalk();
    TalkController.receiveTalk();
});