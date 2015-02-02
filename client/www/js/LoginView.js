define(function (require, exports, module) {
	'use strict';
	var LoginHtml = require('text!/view/login.html'),
		Mustache = require('js/thirdparty/mustache');

	var LoginView = Backbone.View.extend({
		initialize: function (option) {
			document.body.innerHTML = Mustache.render(LoginHtml, {
				strings: global.strings
			});
			this.$el = $(document);
			$(document).foundation();
			$('#signInPanel input[type="text"]')[0].focus();
			this.switch2Mainframe = option || function () {};
		},

		events: {
			'click #signInPanel .button': 'signIn',
			'click #signUpPanel .button': 'signUp'
		},

		signIn: function () {
			var self = this,
				userName = $('#signInPanel input[type="text"]')[0].value.trim(),
				password = $('#signInPanel input[type="password"]')[0].value.trim();
			
			if (!this.validateSignInParams(userName, password)) {
				return;
			}
			
			$.post('/user/login', JSON.stringify({
				name: userName,
				password: password
			})).done(function (user) {
				$('#signInPanel .error').hide();
				global.currentUser = JSON.parse(user);
				self.switch2Mainframe();
			}).fail(function (resp) {
				$('#signInPanel .error').html(resp.responseText);
				$('#signInPanel .error').show();
			});
		},

		signUp: function () {
			var self = this;
			var userName = $('#signUpPanel input[type="text"]')[0].value.trim(),
				password = $('#signUpPanel input[type="password"]')[0].value.trim();
			
			if (!this.validateSignUpParams(userName, password)) {
				return;
			}
			
			// send request to server
			$.post('/user/signup', JSON.stringify({
				name: userName,
				password: password
			})).done(function (user) {
				$('#signUpPanel .error').hide();
				global.currentUser = JSON.parse(user);
				self.switch2Mainframe();
			}).fail(function (resp) {
				$('#signUpPanel .error').html(resp.responseText);
				$('#signUpPanel .error').show();
			});
		},
		
		/**
		 * @return true is valid, otherwise false
		 */
		validateSignInParams: function(userName, password) {
			if (userName.length < 2 || password.length < 6) {
				$('#signInPanel .error').show();
				return false;
			}
			
			return true;
		},
		
		/**
		 * @return true is valid, otherwise false
		 */
		validateSignUpParams: function(userName, password) {
			if (userName.length < 2) {
				$('#signUpPanel .error').html(global.strings.LOGIN_ERR_USER_NAME_TOO_SHORT);
				$('#signUpPanel .error').show();
				return false;
			}
			
			if (password.length < 6) {
				$('#signUpPanel .error').html(global.strings.LOGIN_ERR_PWD_TOO_SHORT);
				$('#signUpPanel .error').show();
				return false;
			}
			
			return true;
		}

	});

	return LoginView;

});