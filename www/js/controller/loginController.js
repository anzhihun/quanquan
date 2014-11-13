define(function (require, exports) {
    'use strict';

    function bindActionHandler() {
        $('#signInPanel .button').click(signIn);
    }

    function signIn() {
        var userName = $('#signInPanel input[type="text"]')[0].value,
            password = $('#signInPanel input[type="password"]')[0].value;
        $.post('/login', JSON.stringify({
            name: userName,
            password: password
        })).done(function(){
            //TODO switch to main window
        }).fail(function(){
            //TODO  show errors
        });
    }

    exports.bindActionHandler = bindActionHandler;
});