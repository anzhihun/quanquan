define(function(require, exports, module){
    'use strict';
    
    var ChannelList = require('js/channel/ChannelList');
    
    var ChannelListView = Backbone.View.extend({
        el: '#channelList',
        initialize: function(){
            // TODO get channel list from server
            this.channels = new ChannelList();
//            this.channels.fetch();
        },
        events: {
            'click #showAddChanDlgBtn': 'openAddChannelDlg'
        },
        
        openAddChannelDlg: function(){
            $('#addChannelDialog').foundation('reveal', 'open');
        }
    });
    
    return ChannelListView;
});