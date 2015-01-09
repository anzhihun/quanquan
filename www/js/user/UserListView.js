define(function(require, exports, module){
	'use strict';
	
	var UserList = require('js/user/UserList');
	var UserItemView = require('js/user/UserItemView');
	
	var UserListView = Backbone.View.extend({
		model: /*@type {UserList}*/new UserList(),
		initialize: function(){
			this.model.bind('change', this.render, this);
			this.model.bind('add', this.addUser, this);
		},
		render: function(){
			var userListContainer = $('.main_frame .right');
			var userItemView = new UserItemView();
			for(var index = 0, len = this.model.length; index < len; index++) {
				userItemView.model = this.model.at(index);
				userListContainer.append(userItemView.render());
			}
		},
		addUser: function(user) {
			var userListContainer = $('.main_frame .right');
			var userItemView = new UserItemView();
			userItemView.model = user;
			userListContainer.append(userItemView.render());
		},
		refresh: function(){
			this.model.url = '/user?channel=global';
			this.model.fetch();
		}
	});

	return UserListView;
});