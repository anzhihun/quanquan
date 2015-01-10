define(function(require, exports, module){
   'use strict';
    
    var TalkMessage = require('js/msg/TalkMessage');
    
    var TalkMessageCollection = Backbone.Collection.extend({
        model: TalkMessage,
		initialize: function(channelName) {
			this.url = '/talk/' + channelName
		}
    });
    
    return TalkMessageCollection;
    
});