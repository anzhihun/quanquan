define(function(require, exports, module){
	'use strict';
	
	var User = Backbone.Model.extend({
		name: '',
		iconUrl: '',
		serverId: 0,
		online: false
	});
	
	return User;
});