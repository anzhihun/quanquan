/* global define */
define(function(require, exports, module){
    function Message() {
        this.MsgType = this.MSG_TYPE_TALK;
        this.To = 'all';
        this.IsToOne = false;
        this.Content = '';
    }
    
    Message.prototype.MSG_TYPE_TALK = 'talk';
    
    return Message;
});