define(function (require, exports, module) {
    'use strict';

    var TalkMessageCollection = require('js/msg/TalkMessageCollection');
    var TalkMessageView = require('js/msg/TalkMessageView');
	
    var MessageBoard = Backbone.View.extend({
        initialize: function (channelName) {
			this.channelName = channelName;
            this.messages = new TalkMessageCollection(channelName);
            this.messages.on('add', this.addMsg, this);
            this.messages.on('reset', this.render, this);
            this.messages.fetch();
            this.$el = $('.main .message_area .body');
        },
        
        render: function(){
        },
        
        getModel: function(){
            return this.messages;
        },
        
        addMsg: function (talkMsg) {
            var msgView = new TalkMessageView(talkMsg);
            if (this.$el.length === 0) {
                this.$el = $('.main .message_area .body');
            }
            this.$el.append(msgView.render());
        },
		
		clear: function(){
			this.stopListening();
			this.off();
			this.undelegateEvents();
		}
    });

    return MessageBoard;

});