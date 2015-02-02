define(function(require, exports, module){
    'use strict';
    
    var ChannelList = require('js/channel/ChannelList'),
		NewChannelDialog = require('js/channel/NewChannelDialog'),
        Mustache = require('js/thirdparty/mustache'),
		ChannelListPanel = require('text!/view/channelListPanel.html'),
        ChannelItemTemplate = require('text!/view/channelItem.html');
	
    var ChannelListView = Backbone.View.extend({
        initialize: function(msgType){
			this._msgType = msgType;
			this._id = msgType + 'ChanList';
			$('.list_container').append(Mustache.render(ChannelListPanel, {
				strings: global.strings,
				id: msgType + 'ChanList'
			}));
			this.$el = $('#' + this._id);
            this.channels = new ChannelList(global.currentUser.name);
			global.currentTalkTarget = {
				name: 'Global',
				isChannel: true
			};
			
			this.channels.on('add', this.renderChannel, this);
			this.channels.on('reset', this.render, this);
			this.channels.fetch({reset: true});
			
			this.newChanDlg = new NewChannelDialog();
        },
		
        events: {
            'click #showAddChanDlgBtn': 'openAddChannelDlg',
			'click li': 'onSelectChannel'
        },
		
		getModel: function() {
			return this.channels;
		},
		
		render: function(){
			var listView = this.$el.find('ul');
			listView.empty();
            this.addDefaultChannels(listView);

			for(var index = 0, len = this.channels.length; index < len; index++) {
                listView.append(Mustache.render(ChannelItemTemplate, {
                    strings: global.strings,
                    chanId: this.channels.at(index).get('name'),
                    chanName: this.channels.at(index).get('name')
                }));
			}
            
            // select default channel 
            this.selectChannel('Global');
		},
        
        addDefaultChannels: function(listView) {
            listView.append(Mustache.render(ChannelItemTemplate, {
                strings: global.strings,
                chanId: 'Global',
                chanName: 'Global'
            }));
        },
		
		renderChannel: function(channel){
			var $list = this.$el.find('ul');
            $list.append(Mustache.render(ChannelItemTemplate, {
                    strings: global.strings,
                    chanId: channel.get('name'),
                    chanName: channel.get('name')
                }));
		},
        
        openAddChannelDlg: function(){
            this.newChanDlg.open();
        },
		
		unselectAll: function() {
			this.$el.find('ul li').each(function(index, elem){
				$(elem).removeClass('active');
			});
		},
		
		onSelectChannel: function(evt){
			var channelId = $(evt.currentTarget).data('chan-id').toString();
			this.selectChannel(channelId);
		},
        
        selectChannel: function(chanId) {
            $('.message_board_toolbar .chan_desc').html(chanId);
            this.unselectAll();
		
			// switch message list
            this.$el.find('[data-chan-id="' + chanId + '"]').addClass('active');
			global.mainframe.switchChannel(chanId);
			
			global.currentTalkTarget = {
				name: chanId,
				isChannel: true
			};
        },
		
		show: function() {
			this.$el.show();
		},
		
		hide: function() {
			this.$el.hide();
		}
    });
    
    return ChannelListView;
});