define(function(require, exports, module){
	'use strict';
	
	var TalkMessage = Backbone.Model.extend({
		user: null,
		content: '',
		dateTime: 0
	});
	
	return TalkMessage;
	
});