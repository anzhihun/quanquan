define(function (require, exports, module) {
	'use strict';
	var MainWindowHtml = require('text!/view/mainWindow.html'),
		UserListView = require('js/user/UserlistView'),
		ChannelListView = require('js/channel/ChannelListView'),
		MessageBoard = require('js/msg/MessageBoard');
  
	var Mainframe = Backbone.View.extend({
		show: function () {
			document.body.innerHTML = MainWindowHtml;
			this.$el = $(document);
			$(document).foundation();
			
			$('#curUser').find('img')[0].src = global.currentUser.iconUrl;
			$('#curUser').find('span')[0].innerHTML = global.currentUser.name;
			
			this.channelListView = new ChannelListView();
			
			this.messageBoard = new MessageBoard();

			this.userListView = new UserListView();
			this.userListView.refresh();
		},
		
		getMessageBoard: function(){
			return this.messageBoard;
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