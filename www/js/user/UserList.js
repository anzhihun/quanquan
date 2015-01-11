define(function(require){
	'use strict';
	
	var User = require('js/user/User');
	
	var UserList = Backbone.Collection.extend({
		model: User,
		url: '/user?channel=Global'
	});
	
	return UserList;
});