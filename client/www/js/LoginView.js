define(function(require, exports, module){
  'use strict';
  var LoginHtml = require('text!/view/login.html'),
      Mustache = require('js/thirdparty/mustache');
  
  var LoginView = Backbone.View.extend({
    initialize: function(option){
      document.body.innerHTML = Mustache.render(LoginHtml, {strings: global.strings});
      this.$el = $(document);
      $(document).foundation();
      $('#signInPanel input[type="text"]')[0].focus();
      this.switch2Mainframe = option || function(){};
    },
    
    events: {
      'click #signInPanel .button': 'signIn',
      'click #signUpPanel .button': 'signUp'
    },
    
    signIn: function(){
      var self = this;
      var userName = $('#signInPanel input[type="text"]')[0].value,
            password = $('#signInPanel input[type="password"]')[0].value;
      $.post('/user/login', JSON.stringify({
          name: userName,
          password: password
      })).done(function(user){
          //TODO switch to main window
          $('#signInPanel .error').hide();
          global.currentUser = JSON.parse(user);
          self.switch2Mainframe();
      }).fail(function(){
          //TODO  show errors
          $('#signInPanel .error').show();
      });
    },
    
    signUp: function(){
      var self = this;
      var userName = $('#signUpPanel input[type="text"]')[0].value,
          password = $('#signUpPanel input[type="password"]')[0].value;
      $.post('/user/signup', JSON.stringify({
          name: userName,
          password: password
      })).done(function(user){
          //TODO switch to main window
          $('#signUpPanel .error').hide();
          global.currentUser = JSON.parse(user);
          self.switch2Mainframe();
      }).fail(function(obj){
          //TODO  show errors
          $('#signUpPanel .error')[0].innerHTML = obj.responseText;
          $('#signUpPanel .error').show();
      });
    }
    
  });
  
  return LoginView;
  
});