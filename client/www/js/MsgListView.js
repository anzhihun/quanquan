define(function (require, exports, module) {
	'use strict';

	var ChannelListView = require('js/channel/ChannelListView'),
		DirectListView = require('js/channel/DirectListView');

	var MsgListView = function (msgType) {
		this._msgType = msgType;
		this._channelListView = null;
		this._directListView = null;

		if (msgType === 'channel') {
			this.loadChannelListView();
		} else if (msgType === 'star') {
			this.loadChannelListView();
			this.loadDirectListView();
		} else if (msgType === 'achived') {
			this.loadChannelListView();
			this.loadDirectListView();
		} else if (msgType === 'direct') {
			this.loadDirectListView();
		}
	};

	MsgListView.prototype.loadChannelListView = function () {
		this._channelListView = new ChannelListView(this._msgType);
	};

	MsgListView.prototype.loadDirectListView = function () {
		this._directListView = new DirectListView(this._msgType);
	};

	MsgListView.prototype.getChannelListView = function () {
		return this._channelListView;
	};

	MsgListView.prototype.getDirectListView = function () {
		return this._directListView;
	};
	
	MsgListView.prototype.show = function() {
		if (this._channelListView) {
			this._channelListView.show();
		} 
		if (this._directListView) {
			this._directListView.show();
		}
	};
	
	MsgListView.prototype.hide = function() {
		if (this._channelListView) {
			this._channelListView.hide();
		} 
		if (this._directListView) {
			this._directListView.hide();
		}
	}

	return MsgListView;
});