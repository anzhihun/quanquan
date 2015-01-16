define(function (require, exports, module) {
    'use strict';

    var TalkMessageCollection = require('js/msg/TalkMessageCollection');
    var TalkMessageView = require('js/msg/TalkMessageView');
	
    var MessageBoard = Backbone.View.extend({
		tagName : 'div',  
        className : 'message_area',  
        initialize: function (channelName) {
			this.channelName = channelName;
            this.messages = new TalkMessageCollection(channelName);
            this.messages.on('add', this.addMsg, this);
            this.messages.on('reset', this.render, this);
            this._isAdded = false;
        },
        
        render: function(){
        },
		
		hide: function() {
			this.$el.css({'display': 'none'});	
		},
		
		show: function(){
			if (this._isAdded === false) {
				this.$el.insertBefore($('.main .send_area'));
				this.$el[0].innerHTML = '<h1>Messages: </h1><div class="body"> </div>';
				this._isAdded = true;
			}
			this.messages.fetch();
			this.$el.css({'display': 'block'});
		},
        
        getModel: function(){
            return this.messages;
        },
        
        addMsg: function (talkMsg) {
            var msgView = new TalkMessageView(talkMsg);
			var msgBody = this.$el.find('.body');
            msgBody.append(msgView.render());
        },
		
		clear: function(){
			this.stopListening();
			this.off();
			this.undelegateEvents();
		}
    });

    return MessageBoard;

});