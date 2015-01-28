define(function(require, exports, module){
	'use strict';
	
	var MessageItemTemplate = require('text!/view/messageItem.html');
	var TalkMessageView = Backbone.View.extend({
        initialize: function(model){
            this.model = model;
        },
        
		render: function(){
			var dateTime = new Date(this.model.get('dataTime'));
			var messageHtml = Mustache.render(MessageItemTemplate, {
                headImg: this.model.get('user').iconUrl,
                name: this.model.get('user').name,
                datetime: dateTime.toLocaleDateString() + ' ' + dateTime.toLocaleTimeString(),
                content: this.model.get('content')
            });
            return messageHtml;
		}
	});
	
	
	return TalkMessageView;
	
});