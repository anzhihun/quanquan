define(function(require, exports, module){
    'use strict';
    
    var ChannelList = require('js/channel/ChannelList');
	var NewChannelDialog = require('js/channel/NewChannelDialog');
	
    var DirectListView = Backbone.View.extend({
        el: '#directList',
        events: {
			'click dd': 'onDialogueClick'
        },
		
		addDialogue: function(userName) {
			var $list = this.$el.find('dl');
			$list.append('<dd><a href="#panel4" style="background: transparent;">' + userName + '</a></dd>');
		},
		
		onDialogueClick: function(evt){
			var userName = $(evt.currentTarget).text();
			this.selectDialogue(userName);
		},
		
		unselectAll: function() {
			this.$el.find('dl dd').each(function(index, elem){
				$(elem).removeClass('active');
				$(elem).find('a').attr('style', 'background: transparent;');
			});
		},
        
		selectDialogue: function(userName){
			var $currentUserElem = null;
			this.$el.find('dl dd').each(function(index, elem){
				$(elem).removeClass('active');
				$(elem).find('a').attr('style', 'background: transparent;');
				if ($(elem).text() === userName) {
					$currentUserElem = $(elem);
				}
			});
			
			// switch message list
			$currentUserElem.addClass('active');
			$currentUserElem.find('a').attr('style', '');

			global.mainframe.switch2DirectDialogue(userName);
			global.currentTalkTarget = {
				name: userName,
				isChannel: false
			};
		}
    });
    
    return DirectListView;
});