package service

import (
	"define"
	"encoding/json"
	"event"
	"net"
	"user"
)

func listenViewMsg() {
	event.On("agent:msg", func(newValue, oldValue interface{}) {
		handleViewMsg(newValue.(map[string]interface{}))
	})

	event.On("user:add", func(newValue, oldValue interface{}) {
		handleAddUserMsg(newValue.(*user.User))
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

func handleAddUserMsg(newUser *user.User) {
	content, err := json.Marshal(newUser)
	if err != nil {
		//TODO log error
		return
	}

	commServer.sendMessage(net.IPv4(255, 255, 255, 255), define.Message{
		MsgType:  define.MSG_TYPE_USER_ADD,
		From:     user.Self.Name,
		HeadImg:  user.Self.HeadImg,
		To:       "all",
		IsPublic: true,
		Content:  string(content),
	})
}
