define(function (require, exports) {
    'use strict';
    var signInCallback = null;
    
    function bindActionHandler(callback) {
        signInCallback = callback || function(){};
        $('#signInPanel .button').click(signIn);
        $('#signUpPanel .button').click(signUp);
    }

    function signIn() {
        var userName = $('#signInPanel input[type="text"]')[0].value,
            password = $('#signInPanel input[type="password"]')[0].value;
        $.post('/user/login', JSON.stringify({
            name: userName,
            password: password
        })).done(function(){
            //TODO switch to main window
            $('#signInPanel .error').hide();
            signInCallback();
        }).fail(function(){
            //TODO  show errors
            $('#signInPanel .error').show();
        });
    }
    
    function signUp() {
        var userName = $('#signUpPanel input[type="text"]')[0].value,
            password = $('#signUpPanel input[type="password"]')[0].value;
        $.post('/user/signup', JSON.stringify({
            name: userName,
            password: password
        })).done(function(){
            //TODO switch to main window
            $('#signUpPanel .error').hide();
            signInCallback();
        }).fail(function(obj){
            //TODO  show errors
            $('#signUpPanel .error')[0].innerHTML = obj.responseText;
            $('#signUpPanel .error').show();
        });
    }

    exports.bindActionHandler = bindActionHandler;
});