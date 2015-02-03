define(function(require, exports, module){
    'use strict';
    
    var ChannelList = require('js/channel/ChannelList'),
		Mustache = require('js/thirdparty/mustache'),
		DirectListPanel = require('text!/view/directListPanel.html'),
		DirectItem = require('text!/view/directItem.html'),
		NewChannelDialog = require('js/channel/NewChannelDialog');
	
    var DirectListView = Backbone.View.extend({
        events: {
			'click li': 'onDialogueClick'
        },
		initialize: function(msgType, container){
			this._msgType = msgType;
			this._id = msgType + 'DirectList';
			$(container).append(Mustache.render(DirectListPanel, {
				strings: global.strings,
				id: this._id,
			}));
			this.$el = $(container).find('#' + this._id);
			this._userNames = [];	
		},
		addDialogue: function(userName) {
			
			if (this._userNames.indexOf(userName) !== -1) {
				return;
			}
			
			var $list = this.$el.find('ul');
			$list.append( Mustache.render(DirectItem, {
				id: userName,
				name: userName,
			}) );
			this._userNames.push(userName);
		},
		
		onDialogueClick: function(evt){
			var userName = $(evt.currentTarget).text();
			this.selectDialogue(userName);
		},
		
		unselectAll: function() {
			this.$el.find('ul li').each(function(index, elem){
				$(elem).removeClass('active');
			});
		},
        
		selectDialogue: function(userName){
			var $currentUserElem = null;
			this.$el.find('ul li').each(function(index, elem){
				$(elem).removeClass('active');
				if ($(elem).text() === userName) {
					$currentUserElem = $(elem);
				}
			});
			
			// switch message list
			$currentUserElem.addClass('active');

			global.mainframe.switch2DirectDialogue(userName);
			global.currentTalkTarget = {
				name: userName,
				isChannel: false
			};
		}, 
		
		getSelectUserId: function() {
			var $userElem = this.$el.find('ul li.active');
			if ($userElem.length === 0) {
				return '';
			} else {
				return $userElem.data('user-id');
			}
		}
    });
    
    return DirectListView;
});