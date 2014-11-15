/*global define, $, alert */
define(function(require, exports, module){
    'use strict';
    
    function getAllUser(callback) {
        $.get('/user?channel=global', function(data, status){
            if (data.indexOf("error:") === 0) {
                alert("failed to get all user: " + data);
            } else {
                callback(JSON.parse(data));
            }
        });
    }
    
    exports.getAllUser = getAllUser;
});