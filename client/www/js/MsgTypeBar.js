define(function (require, exports, module) { 
	'use strict';
	
	var MsgTypeBar = Backbone.View.extend({
		el: '.left-toolbar',
		events: {
			'click li': 'onMsgTypeBtnClick'
		},
		
		onMsgTypeBtnClick: function(evt){
			
			this.$el.find('li').each(function(index, elem){
				$(elem).removeClass('active');
			});
			
			$(evt.currentTarget).addClass('active');
			this.select($(evt.currentTarget).data('id'));
		},
		
		activeBtn: function(id) {
			this.unactiveAllBtn();
			this.$el.find('[data-id="' + id + '"]').addClass('active');
		},
		
		unactiveAllBtn: function() {
			this.$el.find('li').each(function(index, elem){
				$(elem).removeClass('active');
			});
		},
		
		select: function(msgType) {
			this.activeBtn(msgType);
			global.mainframe.switchMsgType(msgType);
		}
		
	});
	
	return MsgTypeBar;
	
	
} );