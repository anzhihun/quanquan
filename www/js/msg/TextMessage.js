define(function(require, exports, module){
	'use strict';
	
	var TextMessage = Backbone.Model.extend({
		user: null,
		content: '',
		dateTime: 0
	})
	
	return TextMessage;
	
});