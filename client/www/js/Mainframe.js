define(function (require, exports, module) {
	'use strict';
	var MainWindowHtml = require('text!/view/mainWindow.html'),
		UserListView = require('js/user/UserlistView'),
//		ChannelListView = require('js/channel/ChannelListView'),
//		DirectListView = require('js/channel/DirectListView'),
		MsgListView = require('js/MsgListView'),
		MessageBoard = require('js/msg/MessageBoard'),
		MessageInputView = require('js/msg/MessageInputView'),
        InviteUserDlg = require('js/user/InviteUserDlg'),
		MsgTypeBar = require('js/msgTypeBar'),
        Mustache = require('js/thirdparty/mustache');
  
	var Mainframe = Backbone.View.extend({
		initialize: function(){
			this._msgBoards = {};
		},
		show: function () {
			document.body.innerHTML = Mustache.render(MainWindowHtml, {strings: global.strings});
			this.$el = $(document);
			$(document).foundation();

			$('#curUser').find('img')[0].src = global.currentUser.iconUrl;
			$('#curUser').find('img')[0].title = global.currentUser.name;
			
			this._msgListView = {};
			
			this.msgListView = null;
			this.msgTypeBar = new MsgTypeBar();
			this.msgTypeBar.select('channel');

			this.userListView = new UserListView('Global');
			this.userListView.refresh();
			
			this.inputView = new MessageInputView();
            
            this.inviteUserDlg = new InviteUserDlg();
            
            $('#inviteUserBtn').click(this.openInviteUserDlg.bind(this));
            $('#showAllUser').click(this.showAllUsers.bind(this));
			
			// select global channel default
			this.switchChannel('Global');
		},
		
		switchMsgType: function(msgType) {
			if (!this._msgListView[msgType]) {
				this._msgListView[msgType] = new MsgListView(msgType);
			}
			
			if (this.msgListView) {
				this.msgListView.hide();
			}
//			
//			for (var key in this._msgListView) {
//				this._msgListView[key].hide();
//			}

			this.msgListView = this._msgListView[msgType];
			this.msgListView.show();
		},
		
		switchChannel: function(channelName) {
			// check
			if (this._msgBoards['chan::'+channelName] === this.messageBoard) {
				return;
			}
			
			$('.message_board_toolbar span').html('#' + channelName);
			
			if (this.msgListView && this.msgListView.getDirectListView()) {
				this.msgListView.getDirectListView().unselectAll();
			}
			
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
			
			if (this.msgListView && this.msgListView.getChannelListView()) {
				this.msgListView.getChannelListView().unselectAll();
			}
			
//			this.channelListView.unselectAll();
			
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
			if (this.msgListView && this.msgListView.getDirectListView()) {
				this.msgListView.getDirectListView().addDialogue(userName);
				this._msgBoards['p2p::'+userName] = new MessageBoard(userName);
			} else if (this._msgListView['direct'] && this._msgListView['direct'].getDirectListView()) {
				this._msgListView['direct'].getDirectListView().addDialogue(userName);
				this._msgBoards['p2p::'+userName] = new MessageBoard(userName);
			} else {
				this._msgListView['direct'] = new MsgListView('direct');
				this._msgListView['direct'].getDirectListView().addDialogue(userName);
				this._msgBoards['p2p::'+userName] = new MessageBoard(userName);
			}
			
//			this.directListView.addDialogue(userName);
//			this._msgBoards['p2p::'+userName] = new MessageBoard(userName);
//			this.directListView.selectDialogue(userName);
		},
		
		clearMessageBoard: function() {
			if (this.messageBoard) {
				this.messageBoard.hide();
			}
			if (!this._msgBoards['blank']) {
				this._msgBoards['blank'] = new MessageBoard('blank');
			}
			this.messageBoard = this._msgBoards['blank'];
			this.messageBoard.show();
			$('.message_board_toolbar span').html('');
		},
        
		openInviteUserDlg: function(evt) {
			this.inviteUserDlg.open();
		},
        
        showAllUsers: function(evt){
            if ($(evt.currentTarget).hasClass('active')) {
                this.userListView.hide();
            } else {
                this.userListView.show();
            }
            $(evt.currentTarget).toggleClass('active');
        },
		
		getDirectListView: function(){
			if (this.msgListView) {
				return this.msgListView.getDirectListView();
			} else {
				return null;
			}
//			return this.directListView;	
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
			if (this.msgListView) {
				return this.msgListView.getChannelListView();
			} else {
				return null;
			}
//			return this.channelListView;
		}
										
	});

	return Mainframe;
});