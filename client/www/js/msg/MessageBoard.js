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
			
//			this.$el.css({'display': 'none'});
			$('.main .msg_board_container').append(this.$el);
//			this.$el.insertAfter($('.main .message_board_toolbar'));
            this.$el[0].innerHTML = '<div class="body"> </div>'
			
			this.messages.fetch();
        },
        
        render: function(){
        },
		
		hide: function() {
			this.$el.removeClass('left-to-right-hide');
			this.$el.removeClass('left-to-right-show');
			this.$el.addClass('left-to-right-hide');
		},
		
		show: function(){
			this.$el.removeClass('left-to-right-hide');
			this.$el.removeClass('left-to-right-show');
			this.$el.addClass('left-to-right-show');
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