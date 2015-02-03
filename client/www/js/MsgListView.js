define(function (require, exports, module) {
	'use strict';

	var ChannelListView = require('js/channel/ChannelListView'),
		DirectListView = require('js/channel/DirectListView');

	var MsgListView = function (msgType) {
		
		$('.left-title').append('<div class="list_container" id="list-' + msgType  + '"></div>');
		this.$el = $('.left-title #list-' + msgType);
		
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
		this._channelListView = new ChannelListView(this._msgType, this.$el);
	};

	MsgListView.prototype.loadDirectListView = function () {
		this._directListView = new DirectListView(this._msgType, this.$el);
	};

	MsgListView.prototype.getChannelListView = function () {
		return this._channelListView;
	};

	MsgListView.prototype.getDirectListView = function () {
		return this._directListView;
	};
	
	MsgListView.prototype.show = function() {
		this.$el.removeClass('left-to-right-hide');
		this.$el.removeClass('left-to-right-show');
		this.$el.addClass('left-to-right-show');

		if (this._channelListView) {
			var chanId = this._channelListView.getSelectChannelId();
			if (chanId.length !== 0) {
				this._channelListView.selectChannel(chanId);
				return;
			}
		} 
		
		if (this._directListView) {
			var userId = this._directListView.getSelectUserId();
			if (userId.length !== 0) {
				this._directListView.selectDialogue(userId);
				return;
			}
		}
		
		// clear message board
		global.mainframe.clearMessageBoard();
		
	};
	
	MsgListView.prototype.hide = function() {
		this.$el.removeClass('left-to-right-hide');
		this.$el.removeClass('left-to-right-show');
		this.$el.addClass('left-to-right-hide');
	}

	return MsgListView;
});