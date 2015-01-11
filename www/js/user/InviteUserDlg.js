define(function(require, exports, module){
	'use strict';
	
	var InviteUserDlgTemplate = require('text!/view/inviteUserDlg.html');
	var UserList = require('js/user/UserList');
	
	var InviteUserDlg = Backbone.View.extend({
		el: '#inviteUserDialog',
		initialize: function(){
			this.model = new UserList();
			this.model.on('reset', this.render, this);
		},
		open: function(){
			var self = this;
			this.model.fetch({reset: true, success: function(){
				self.$el.foundation('reveal', 'open');
			}});
		},
		
		events: {
			'click .button': 'inviteUsers'
		},
		
		inviteUsers: function(){
			//TODO 
		},
		
		render: function() {
			this.$el.empty();
			var users = [];
			var userModel = this.model;
			for(var index = 0, len = userModel.length; index < len; index++) {
				users.push({
					name: userModel.at(index).get('name'),
					iconUrl: userModel.at(index).get('iconUrl')
				});
			}
			var dlgHtml = Mustache.render(InviteUserDlgTemplate, {
				users: users
			});
			this.$el.append(dlgHtml);
		},
		
		close: function() {
			this.$el.foundation('reveal', 'close');
			this.$el.empty();
		}						
	});
	
	return InviteUserDlg;

});