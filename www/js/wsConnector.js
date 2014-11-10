define(function () {
    'use strict';

    function WSConnector(url, messageHandler, callbackOnOpen) {
        this._url = url;
        this._socket = null;
        this._timeOutSecond = 1000;
        this._messageHandler = messageHandler;
        this._callbackOnOpen = callbackOnOpen || function(){};
        this._isConnected = false;

        this.connect();
    }

    WSConnector.prototype.connect = function () {
        
        var self = this;
        
        if (self._isConnected) {
            return;
        }
        
        this._socket = new WebSocket(this._url);

        this._socket.onopen = function () {
            console.log('CONNECTED');
            self._isConnected = true;
            self.setTimeOutSecond(1000);
            self._callbackOnOpen(self);
        };

        this._socket.onclose = function () {
            console.log('DISCONNECTED');
            self._isConnected = false;
            self.reConnect();
        };

        this._socket.onmessage = function (evt) {
//            console.log('RESPONSE: ' + evt.data);
//            self.parseMsg(evt.data);
            self._messageHandler.onMessage(evt.data);
        };

        this._socket.onerror = function (evt) {
            console.log('ERROR: ' + evt.data);
            self._isConnected = false;
            self.reConnect();
        };
    };

    WSConnector.prototype.sendMessage = function (msg) {
        console.log('send msg to server: ' + msg);
        this._socket.send(msg);
    };

    WSConnector.prototype.reConnect = function () {
    
        this._timeOutSecond = this._timeOutSecond * 2;
        var self = this;
        
        setTimeout(function(){
            self.connect();
        },self.getTimeOutSecond());
    };
    
    WSConnector.prototype.getTimeOutSecond = function(){
        
        return this._timeOutSecond;
    };
    
    WSConnector.prototype.setTimeOutSecond = function(second){
        this._timeOutSecond = second;
    };

    return WSConnector;
});