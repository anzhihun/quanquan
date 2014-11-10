package controller

import (
	"code.google.com/p/go.net/websocket"
	"event"
	"fmt"
	"github.com/gocraft/web"
	"utils"
)

type rtMessageController struct {
	conn *websocket.Conn
}

var rtMsgController rtMessageController

func WebSocketMiddleware(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	if req.URL.Path == "/rtmsg" {
		websocket.Handler(rtMsgController.listenRTConnect).ServeHTTP(rw, req.Request)
	} else {
		next(rw, req)
	}
}

// 处理每一个websocket请求
func (this *rtMessageController) listenRTConnect(ws *websocket.Conn) {
	this.conn = ws
	for {
		var recvMsg string
		if err := websocket.Message.Receive(ws, &recvMsg); err != nil {
			return
		}

		if msgData, err := utils.DecodeJsonMsg(recvMsg); err != nil { //解析json字符串
			// reportRtMsgError(errors.RtMsgError{self.connection.RemoteAddr().String(), self.clientId, "ultils.DecodeJsonMsg", err.Error(), "无法处理消息，消息解析失败！"}, self.connection)
			continue
		} else if msgType, ok := msgData["MsgType"].(string); !ok { //解析消息类型
			// reportRtMsgError(errors.RtMsgError{self.connection.RemoteAddr().String(), self.clientId, "parse recvMsg", "can not parse msg type", "无法处理消息，消息类型解析失败！"}, self.connection)
		} else { //根据消息类型进一步解析消息内容
			this.handMessage(msgType, msgData)
		}
	}
}

func (this *rtMessageController) listenToViewMessage() {
	event.On("view:msg", func(newValue, oldValue interface{}) {
		msg := newValue.([]byte)
		this.sendMessage(msg)
	})
}

func (this *rtMessageController) handMessage(msgType string, msgData map[string]interface{}) {
	event.Trigger("agent:msg", msgData, nil)
}

func (this *rtMessageController) sendMessage(msg []byte) {
	if err := websocket.Message.Send(this.conn, string(msg)); err != nil {
		// log.GetLogger().Warning(fmt.Sprintf("rtMsg Send failed\n remote address:%s\n clientId:%d\n message:%s\n error:%s\n", self.connection.RemoteAddr().String(), self.clientId, msg, err.Error()))
	}
}
