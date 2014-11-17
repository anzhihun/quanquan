package service

import (
	"conn"
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

	event.On(event.EVENT_F2B_ADD_USER, func(newValue, oldValue interface{}) {
		handleF2BAddUserMsg(newValue.(define.AddUserMessage))
	})

	event.On(event.EVENT_B2F_ADD_USER, func(newValue, oldValue interface{}) {
		handleB2FAddUserMsg(newValue.(*user.User))
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

func handleF2BAddUserMsg(newUser define.AddUserMessage) {

	msgContent, err := json.Marshal(newUser)
	if err != nil {
		// TODO log error
		return
	}

	// broadcast to all agent, include self
	commServer.sendMessage(net.IPv4(255, 255, 255, 255), define.Message{
		MsgType:  define.MSG_TYPE_USER_ADD,
		From:     "",
		HeadImg:  "",
		To:       "all",
		IsPublic: true,
		Content:  string(msgContent),
	})
}

func handleB2FAddUserMsg(newUser *user.User) {

	userContent, err := json.Marshal(newUser)
	if err != nil {
		// TODO log error
		return
	}

	var msgContent []byte
	msgContent, err = json.Marshal(define.ToClientMessage{MsgType: define.MSG_TYPE_USER_ADD, Content: string(userContent)})
	if err != nil {
		// TODO log error
		return
	}

	// broadcast to all websocket which is connect to me
	conn.Broadcast2AllClient(msgContent)
}
