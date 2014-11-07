package controller

import (
	"code.google.com/p/go.net/websocket"
	"net/http"
)

func Init() {
	http.Handle("/rtmsg", websocket.Handler(rtMsgController.listenRTConnect))
	rtMsgController.listenToViewMessage()
}
