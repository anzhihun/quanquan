define(function(require, exports, module){
    'use strict';
    
    var Channel = Backbone.Model.extend({
        defaults: function(){
            return {
                name: 'new channel',
                id: 0
            };
        }
    });
    
    return Channel;
});