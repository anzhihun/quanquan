/* global define, $ */
define(function(require, exports, module){
    'use strict';
    
    function addNewChannel(name, callback) {
        callback = callback || function(){};
        
        $.post('/channel', JSON.stringify({name:name}), function(data, status){
           callback(data); 
        });
    }
    
    exports.addNewChannel = addNewChannel;
});