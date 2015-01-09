define(function(require, exports){
    'use strict';
    
//    var UserView = require('js/view/userView');
    
    // msg handlers 
    var handlers = [{
        msgType: 'addUser',
        handle: function(msg) {
            var user = JSON.parse(msg.Content);
//            UserView.addUser(user);
        }
    },{
		msgType: 'userLogin',
        handle: function(msg) {
            var user = JSON.parse(msg.Content);
//            UserView.login(user);
        }
	}];
    
    function handle(msg) {

        for (var index = 0, len = handlers.length; index < len; index++) {
            if (handlers[index].msgType === msg.MsgType) {
                handlers[index].handle(msg);
                return;
            }
        }
    }
    
    function accept(msg) {
        if (msg === null || msg === undefined ) {
            return false;
        }
        
        for (var index = 0, len = handlers.length; index < len; index++) {
            if (handlers[index].msgType === msg.MsgType) {
                return true;
            }
        }
        
        return false;
    }

    exports.handle = handle;
    exports.accept = accept;
});