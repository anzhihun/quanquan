define(function(require, exports, module){
   'use strict';
    
    var TalkMessage = require('js/msg/TalkMessage');
    
    var TalkMessageCollection = Backbone.Collection.extend({
        model: TalkMessage,
        url: '/talk'
    });
    
    return TalkMessageCollection;
    
});