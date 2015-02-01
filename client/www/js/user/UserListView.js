define(function (require, exports, module) {
	'use strict';

	var UserList = require('js/user/UserList'),
        UserPanel = require('text!/view/userPanel.html'),
        UserItemView = require('js/user/UserItemView'),
        Mustache = require('js/thirdparty/mustache');

	var UserListView = Backbone.View.extend({
		el: '.main_frame .user_list',
		
		model: /*@type {UserList}*/ new UserList(),
		initialize: function (channelName) {
			this.channelName = channelName;
			this.model.bind('change', this.render, this);
			this.model.bind('add', this.addUser, this);
			this.model.bind('reset', this.render, this);
		},
		render: function () {
			this.$el.empty();
			this.$el.append(Mustache.render(UserPanel, {strings: global.strings}));
			
			var userItemView = new UserItemView();
			for (var index = 0, len = this.model.length; index < len; index++) {
				userItemView.model = this.model.at(index);
				this.$el.append(userItemView.render());
			}
			
			if (this.channelName === 'Global') {
				global.allUsers = this.model;
			}
		},
        show: function(evt) {
            this.$el.css({'display':'block'});
        },
        hide: function(evt) {
            this.$el.css({'display':'none'});
        },
		events: {
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
		clear: function(){
			this.stopListening();
			this.off();
			this.undelegateEvents();
			this.$el.empty();
		}
	});

	return UserListView;
});