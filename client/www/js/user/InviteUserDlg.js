define(function(require, exports, module){
	'use strict';
	
	var InviteUserDlgTemplate = require('text!/view/inviteUserDlg.html');
	var UserList = require('js/user/UserList');
	
	var InviteUserDlg = Backbone.View.extend({
		el: '#inviteUserDialog',
		initialize: function(){
			this.clear();
			this.model = new UserList();
			this.model.on('reset', this.render, this);
		},
		open: function(){
			var self = this;
			this.model.fetch({reset: true, success: function(){
				self.$el.foundation('reveal', 'open');
			}});
		},
		
		inviteUsers: function(){
			var inviteUserNames = [];
			var userName = '';
			this.$el.find('input').each(function(index, elem){
				if (elem.checked) {
					userName = $(elem).parent().find('a').text().trim();
					inviteUserNames.push(userName);
				}
			});
			
			// send request 
			var self = this;
			$.post('/channel/inviteUser', JSON.stringify({
			  userNames: inviteUserNames,
			  channelName: global.currentTalkTarget.name,
			  inviter: global.currentUser.name
			})).done(function(user){
			  //TODO switch to main window
				self.close();
			}).fail(function(){
			  //TODO  show errors
				self.close();
			});
		},
		
		render: function() {
			this.clear();
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
			this.$el.find('.button').click(this.inviteUsers.bind(this));
		},
		
		close: function() {
			this.$el.foundation('reveal', 'close');
			this.$el.empty();
		},
		
		clear: function() {
			this.$el.empty();
			this.stopListening();
			this.off();
			this.undelegateEvents();
		}
	});
	
	return InviteUserDlg;

});