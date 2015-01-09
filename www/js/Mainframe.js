define(function (require, exports, module) {
    'use strict';
    var MainWindowHtml = require('text!/view/mainWindow.html'),
        UserListView = require('js/user/UserlistView'),
        ChannelListView = require('js/channel/ChannelListView');

    var Mainframe = Backbone.View.extend({
        show: function () {
            document.body.innerHTML = MainWindowHtml;
            this.$el = $(document);
            $(document).foundation();
            
            this.userListView = new UserListView(); 
            this.userListView.refresh();
            
            this.channelListView = new ChannelListView();
        }
    });

    return Mainframe;
});