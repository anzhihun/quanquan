/* global define, require, document */

require.config({
    baseUrl: '/',
    paths: {
        text: 'js/thirdparty/text'
    }
});


define( function(require, exports, module) {
'use strict';
    
    var MainWindowHtml = require('text!../view/mainWindow.html');
    document.body.innerHTML = MainWindowHtml;
} );