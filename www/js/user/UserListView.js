define(function (require, exports, module) {
	'use strict';

	var UserList = require('js/user/UserList');
	var UserItemView = require('js/user/UserItemView');

	var UserListView = Backbone.View.extend({
		model: /*@type {UserList}*/ new UserList(),
		initialize: function (channelName) {
			this.channelName = channelName;
			this.model.bind('change', this.render, this);
			this.model.bind('add', this.addUser, this);
			this.model.bind('reset', this.render, this);
		},
		render: function () {
			var userListContainer = $('.main_frame .right');
			userListContainer.empty();
			userListContainer.append('<h4>users</h4>');
			var userItemView = new UserItemView();
			for (var index = 0, len = this.model.length; index < len; index++) {
				userItemView.model = this.model.at(index);
				userListContainer.append(userItemView.render());
			}
		},
		addUser: function (user) {
			var userListContainer = $('.main_frame .right');
			var userItemView = new UserItemView();
			userItemView.model = user;
			userListContainer.append(userItemView.render());
		},
		refresh: function () {
			this.model.url = '/user?channel=' + this.channelName;
			this.model.fetch({reset: true});
		},
		getUsers: function () {
			return this.model;
		},
		
		clear: function(){
			this.stopListening();
			this.off();
			this.undelegateEvents();
		}
	});

	return UserListView;
});