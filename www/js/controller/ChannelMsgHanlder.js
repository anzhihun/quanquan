define( function(require, exports, module){
	'use strict';

	var messageBoard = null;
    // msg handlers 
    var handlers = [{
        msgType: 'newChannel',
        handle: function(msg) {
			global.mainframe.getChannelListView().getModel().add(msg.channel);
        }
    }];
	
    function handle(msg) {
		
		messageBoard = messageBoard || global.mainframe.getMessageBoard();
		
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