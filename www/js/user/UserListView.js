define(function (require, exports, module) {
	'use strict';

	var UserList = require('js/user/UserList');
	var UserItemView = require('js/user/UserItemView');
	var InviteUserDlg = require('js/user/InviteUserDlg');

	var UserListView = Backbone.View.extend({
		el: '.main_frame .right',
		
		model: /*@type {UserList}*/ new UserList(),
		initialize: function (channelName) {
			this.channelName = channelName;
			this.model.bind('change', this.render, this);
			this.model.bind('add', this.addUser, this);
			this.model.bind('reset', this.render, this);
		},
		render: function () {
			this.$el.empty();
			this.$el.append('<span>users</span><a id="inviteUsers" style="position: relative;float: right;margin-right: 1em;color: white;cursor: pointer;">+</a>');
			
			var userItemView = new UserItemView();
			for (var index = 0, len = this.model.length; index < len; index++) {
				userItemView.model = this.model.at(index);
				this.$el.append(userItemView.render());
			}
			
			this.inviteUserDlg = new InviteUserDlg();
		},
		events: {
			'click #inviteUsers': 'openInviteUserDlg',
			'click .talk_button': 'onTalkButtonClick'
		},
		onTalkButtonClick: function(evt){
			var userName = $(evt.currentTarget).parent().find('.user_name').text().trim();
			console.log('on onTalkButtonClick userName: ' + userName);
			global.mainframe.addDirectDialogue(userName);
			global.mainframe.getDirectListView().selectDialogue(userName);
		},
		addUser: function (user) {
			var userItemView = new UserItemView();
			userItemView.model = user;
			this.$el.append(userItemView.render());
		},
		refresh: function () {
			this.model.url = '/user?channel=' + this.channelName;
			this.model.fetch({reset: true});
		},
		getUsers: function () {
			return this.model;
		},
		
		openInviteUserDlg: function() {
			this.inviteUserDlg.open();
		},
		
		clear: function(){
			this.inviteUserDlg.clear();
			this.stopListening();
			this.off();
			this.undelegateEvents();
			this.$el.empty();
		}
	});

	return UserListView;
});