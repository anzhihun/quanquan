define(function(require, exports, module){
   'use strict';
    
    var Channel = require('js/channel/Channel');
    
    var ChannelList = Backbone.Collection.extend({
		initialize: function(userName){
			this.url = '/channel?user=' + userName;
		},
        model: Channel
    });
    
    return ChannelList;
});