define(function(require, exports, module){
   'use strict';
    
    var Channel = require('js/channel/Channel');
    
    var ChannelList = Backbone.Collection.extend({
        url: '',
        model: Channel
    });
    
    return ChannelList;
});