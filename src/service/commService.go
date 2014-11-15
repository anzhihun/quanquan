package service

import (
	"define"
	"encoding/json"
	"event"
	"fmt"
	"net"
	"user"
	"utils"
)

type CommunicationService struct {
}

var commServer CommunicationService

// communicate with other quanquan clients throgh udp protocol
func (this *CommunicationService) start() {
	fmt.Println("start server")
	socket, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 53241,
	})

	if err != nil {
		fmt.Println("failed to listen on port 53421!", err)
		return
	}
	defer socket.Close()

	for {
		// read data
		data := make([]byte, 4096)
		read, remoteAddr, err := socket.ReadFromUDP(data)
		fmt.Println("receive msg !")
		if err != nil {
			fmt.Println("Failed to read data from others!", err)
			continue
		}
		go this.handleMessage(remoteAddr.IP, data[:read])
	}
}

func (this *CommunicationService) handleMessage(remoteIp net.IP, msg []byte) {
	// fmt.Printf("receive msg: %s\n\n", msg)
	msgMap, _ := utils.DecodeJsonMsg(string(msg))
	if msgMap["MsgType"].(string) == define.MSG_TYPE_ONLINE {
		// update online status
		user.Online(msgMap["Content"].(string))
		event.Trigger("view:msg", msg, nil)

		// response
		if user.Self != nil {
			this.sendMessage(remoteIp, define.Message{
				MsgType:  define.MSG_TYPE_ACK_ONLINE,
				From:     user.Self.Name,
				HeadImg:  user.Self.HeadImg,
				To:       msgMap["Content"].(string),
				IsPublic: true,
				Content:  "",
			})
		}

	} else if msgMap["MsgType"].(string) == define.MSG_TYPE_ACK_ONLINE {
		// newUser := user.User{msgMap["From"].(string), remoteIp.String(), msgMap["HeadImg"].(string), true}
		// user.AddUser(&newUser)
		event.Trigger("view:msg", msg, nil)

	} else if msgMap["MsgType"].(string) == define.MSG_TYPE_OFFLINE {
		// remove user
		// newUser := user.User{msgMap["From"].(string), remoteIp.String(), msgMap["HeadImg"].(string), false}
		// user.RemoveUser(&newUser)
		event.Trigger("view:msg", msg, nil)

	} else if msgMap["MsgType"].(string) == define.MSG_TYPE_JOIN {
		event.Trigger("view:msg", msg, nil)

	} else if msgMap["MsgType"].(string) == define.MSG_TYPE_TALK {
		event.Trigger("view:msg", msg, nil)

	} else {
		fmt.Println("unknown msg: ", msgMap)
	}
}

func (this *CommunicationService) broadcastMe() {
	if user.Self != nil {
		this.sendMessage(net.IPv4(255, 255, 255, 255), define.Message{
			MsgType:  define.MSG_TYPE_ONLINE,
			From:     user.Self.Name,
			HeadImg:  user.Self.HeadImg,
			To:       "all",
			IsPublic: true,
			Content:  "",
		})
	}
}

func (this *CommunicationService) sendMessage(remoteIp net.IP, msg interface{}) error {
	socket, err := net.DialUDP("udp4", nil, &net.UDPAddr{
		IP:   remoteIp,
		Port: 53241,
	})
	if err != nil {
		fmt.Println("failed to dial udp!", err)
		return err
	}
	defer socket.Close()

	// 发送数据
	senddata, _ := json.Marshal(msg)
	_, err = socket.Write(senddata)
	if err != nil {
		fmt.Println("failed send msg: ", senddata, err)
		return err
	}

	return nil
}
