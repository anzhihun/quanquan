define(function(require, exports){
    'use strict';
    
//    var UserView = require('js/view/userView');
    
    // msg handlers 
    var handlers = [{
        msgType: 'newUser',
        handle: function(msg) {
			global.mainframe.getUserListView().getUsers().add(msg.user);
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
            if (handlers[index].msgType === msg.msgType) {
                handlers[index].handle(msg);
                return true;
            }
        }
		
		return false;
    }

    exports.handle = handle;
});