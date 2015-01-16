define(function (require, exports, module) {
	'use strict';
	var MainWindowHtml = require('text!/view/mainWindow.html'),
		UserListView = require('js/user/UserlistView'),
		ChannelListView = require('js/channel/ChannelListView'),
		DirectListView = require('js/channel/DirectListView'),
		MessageBoard = require('js/msg/MessageBoard'),
		MessageInputView = require('js/msg/MessageInputView');
  
	var Mainframe = Backbone.View.extend({
		initialize: function(){
			this._msgBoards = {};
		},
		show: function () {
			document.body.innerHTML = MainWindowHtml;
			this.$el = $(document);
			$(document).foundation();

			$('#curUser').find('img')[0].src = global.currentUser.iconUrl;
			$('#curUser').find('span')[0].innerHTML = global.currentUser.name;
			
			this.channelListView = new ChannelListView();
			this.messageBoard = new MessageBoard('Global');
			this._msgBoards['chan::Global'] = this.messageBoard;
			this.messageBoard.show();
			
			this.directListView = new DirectListView();

			this.userListView = new UserListView('Global');
			this.userListView.refresh();
			
			this.inputView = new MessageInputView();
		},
		
		switchChannel: function(channelName) {
			// check
			if (this._msgBoards['chan::'+channelName] === this.messageBoard) {
				return;
			}
			
			this.directListView.unselectAll();
			
			// change message board and user list
			if (this.userListView) {
				this.userListView.clear();
			}
			
			this.messageBoard.hide();
			if (this._msgBoards['chan::'+channelName]) {
				this.messageBoard = this._msgBoards['chan::'+channelName];
			} else {
				this.messageBoard = new MessageBoard(channelName);
				// cache
				this._msgBoards['chan::'+channelName] = this.messageBoard;
			}
			this.messageBoard.show();
			
			this.userListView = new UserListView(channelName);
			this.userListView.refresh();
		},
		
		switch2DirectDialogue: function(userName) {
			// check
			var boardId = 'p2p::'+userName;
			if (this._msgBoards[boardId] === this.messageBoard) {
				return;
			}
			
			this.channelListView.unselectAll();
			
			// change message board and user list
			if (this.userListView) {
				this.userListView.clear();
				this.userListView = null;
			}
			
			this.messageBoard.hide();
			if (this._msgBoards[boardId]) {
				this.messageBoard = this._msgBoards[boardId];
			} else {
				this.messageBoard = new MessageBoard(userName);
				// cache
				this._msgBoards[boardId] = this.messageBoard;
			}
			this.messageBoard.show();
		},
		
		addDirectDialogue: function(userName){
			this.directListView.addDialogue(userName);
			this._msgBoards['p2p::'+userName] = new MessageBoard(userName);
//			this.directListView.selectDialogue(userName);
		},
		
		getMessageBoard: function(boardId){
			return this._msgBoards[boardId];
		},
		
		getCurrentMessageBoard: function(){
			return this._msgBoards['chan::'+global.currentTalkTarget.name];
		},
		
		getUserListView: function(){
			return this.userListView;
		},
		
		getChannelListView: function() {
			return this.channelListView;
		}
										
	});

	return Mainframe;
});