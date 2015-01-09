define(function(require, exports, module){
	'use strict';
	
	var TextMessage = require('js/msg/TextMessage');
	var MessageItemTemplate = require('text!/view/messageItem.html');
	
	var TextMessageView = Backbone.Model.extend({
		model: new TextMessage(),
		render: function(){
			var messageContainer = $('.main .message_area .body');
			var dateTime = new Date(this.model.dataTime);
			var messageHtml = Mustache.render(MessageItemTemplate, {
                headImg: this.model.get('user').iconUrl
                name: this.model.get('user').name,
                datetime: dateTime.toLocaleDateString() + ' ' + dateTime.toLocaleTimeString(),
                content: this.model.get('content')
            }));
			messageContainer.append(messageHtml);
			
		}
	});
	
	
	return TextMessageView;
	
});