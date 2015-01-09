define(function(require, exports, module){
	'use strict';
	
	var UserListItemTemplate = require('text!/view/userListItem.html');
	
	var UserItemView = Backbone.View.extend({
		model: /*@type {User}*/null,
		render: function(){
			return Mustache.render(UserListItemTemplate, {
                headImg: this.model.get('iconUrl'),
                name: this.model.get('name')
            });
		}
	})
	
	return UserItemView;
	
});