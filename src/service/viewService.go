package service

import (
	"define"
	"event"
)

func listenViewMsg() {
	event.On("agent:msg", func(newValue, oldValue interface{}) {
		handleViewMsg(newValue.(map[string]interface{}))

	})
}

func handleViewMsg(msgMap map[string]interface{}) {
	msgType := msgMap["MsgType"].(string)
	if msgType == define.MSG_TYPE_TALK {
		commServer.sendMessage(net.IPv4(255, 255, 255, 255), define.Message{
			MsgType:  define.MSG_TYPE_TALK,
			From:     user.Self.Name,
			HeadImg:  user.Self.HeadImg,
			To:       msgMap["To"].(string),
			IsPublic: true,
			Content:  msgMap["Content"].(string),
		})
	}
}
