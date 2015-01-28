define(function (require) {
	'use strict';

	var TalkController = require('js/controller/talkController'),
		UserController = require('js/controller/userController'),
		ChannelMsgHanlder = require('js/controller/ChannelMsgHanlder');

	function WSMsgHandler() {}

	// 接收实时消息
	WSMsgHandler.prototype.onMessage = function (msg) {

		console.log('receive server msg: ' + msg);
		if (msg === null || msg.trim().length === 0) {
			return;
		}

		var msgObj = JSON.parse(msg);
		if (TalkController.handle(msgObj)) {
			return;
		} else if (UserController.handle(msgObj)) {
			return;
		} else if (ChannelMsgHanlder.handle(msgObj)) {
			return;
		} else {
			console.warn('未知消息' + msg);
		}

	};

	return WSMsgHandler;
});