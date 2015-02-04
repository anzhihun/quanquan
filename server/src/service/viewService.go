package service

import (
	"conn"
	"define"
	"encoding/json"
	"entity"
	"event"
	"fmt"
	"net"
)

func listenViewMsg() {
	event.On("agent:msg", func(newValue, oldValue interface{}) {
		handleViewMsg(newValue.(map[string]interface{}))
	})

	event.On(event.EVENT_F2B_ADD_USER, func(newValue, oldValue interface{}) {
		handleF2BAddUserMsg(newValue.(define.AddUserMessage))
	})

	event.On(event.EVENT_B2F_ADD_USER, func(newValue, oldValue interface{}) {
		handleB2FAddUserMsg(newValue.(*entity.User))
	})

	event.On(event.EVENT_F2B_LOGIN, func(newValue, oldValue interface{}) {
		handleF2BUserLogin(newValue.(string))
	})
	event.On(event.EVENT_B2F_LOGIN, func(newValue, oldValue interface{}) {
		handleB2FUserLogin(newValue.(string))
	})

}

func handleViewMsg(msgMap map[string]interface{}) {
	msgType := msgMap["MsgType"].(string)
	fmt.Println("handle view msg from client: ", msgType)
	if msgType == define.MSG_TYPE_TALK {
		commServer.sendMessage(net.IPv4(255, 255, 255, 255), define.Message{
			MsgType:  define.MSG_TYPE_TALK,
			From:     Self.Name,
			HeadImg:  Self.HeadImg,
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

func handleB2FAddUserMsg(newUser *entity.User) {

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

func handleF2BUserLogin(userName string) {
	// broadcast to all agent, include self
	commServer.sendMessage(net.IPv4(255, 255, 255, 255), define.Message{
		MsgType:  define.MSG_TYPE_USER_LOGIN,
		From:     "",
		HeadImg:  "",
		To:       "all",
		IsPublic: true,
		Content:  userName,
	})
}

func handleB2FUserLogin(userName string) {
	// broadcast to all agent, include self
	fmt.Println("user login: ", userName)
	loginUser := FindUser(userName)
	if loginUser == nil {
		// user not exist
		// log
		return
	}

	userContent, err := json.Marshal(loginUser)
	if err != nil {
		//TODO log error
		fmt.Errorf("json marshal user error on user login! %s", err.Error())
		return
	}

	var msgContent []byte
	msgContent, err = json.Marshal(define.ToClientMessage{
		MsgType: define.MSG_TYPE_USER_LOGIN,
		Content: string(userContent),
	})
	if err != nil {
		// TODO log error
		return
	}

	// broadcast to all websocket which is connect to me
	conn.Broadcast2AllClient(msgContent)
}
