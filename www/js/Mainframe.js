define(function(require, exports, module){
	'use strict';
	var MainWindowHtml = require('text!/view/mainWindow.html'),
		UserListView = require('js/user/UserlistView');
	
	var userListView = new UserListView();

	var Mainframe = Backbone.View.extend({
		show: function(){
			document.body.innerHTML = MainWindowHtml;
			$(document).foundation();
			
			userListView.refresh();
		}
	});
	
	return Mainframe;
});