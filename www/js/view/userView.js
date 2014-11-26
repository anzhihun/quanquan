/* global define, $, Mustache*/
define(function(require, exports){
   'use strict';
    
    var UserListItemTemplate = require('text!/view/userListItem.html');
    
    function updateAllUser(users) {
        var userContainer = $('.main_frame .right');
        users.forEach(function(user){
            userContainer.append(Mustache.render(UserListItemTemplate, {
                headImg: user.HeadImg,
                name: user.Name
            }));
        });
    }
    
    function addUser(user) {
        var userContainer = $('.main_frame .right');
        userContainer.append(Mustache.render(UserListItemTemplate, {
            headImg: user.HeadImg,
            name: user.Name
        }));
    }
	
	function login(user) {
		//TODO modify user status
		console.log('user login: ' + user.Name);
	}
    
    exports.updateAllUser = updateAllUser;
    exports.addUser = addUser;
	exports.login = login;
});