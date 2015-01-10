define(function(require, exports, module){
    'use strict';
    
    var ChannelList = require('js/channel/ChannelList');
	var NewChannelDialog = require('js/channel/NewChannelDialog');
	
    var ChannelListView = Backbone.View.extend({
        el: '#channelList',
        initialize: function(){
            // TODO get channel list from server
            this.channels = new ChannelList();
			global.currentTalkTarget = {
				name: 'global',
				isChannel: true
			};
			
			this.channels.on('add', this.renderChannel, this);
			
			this.newChanDlg = new NewChannelDialog();
//            this.channels.fetch();
        },
        events: {
            'click #showAddChanDlgBtn': 'openAddChannelDlg'
        },
		
		getModel: function() {
			return this.channels;
		},
		
		renderChannel: function(channel){
			var $list = this.$el.find('dl');
			$list.append('<dd><a href="#panel4" style="background: transparent;">' + channel.get('name') + '</a></dd>');
		},
        
        openAddChannelDlg: function(){
            this.newChanDlg.open();
        }
    });
    
    return ChannelListView;
});