package server

import (
	"define"
	"encoding/json"
	"event"
	"fmt"
	"net"
	"user"
	"utils"
)

type CommunicationServer struct {
}

var commServer CommunicationServer

// communicate with other quanquan clients throgh udp protocol
func (this *CommunicationServer) start() {
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

func (this *CommunicationServer) handleMessage(remoteIp net.IP, msg []byte) {
	fmt.Printf("receive msg: %s\n\n", msg)
	msgMap, _ := utils.DecodeJsonMsg(string(msg))
	if msgMap["MsgType"].(string) == define.MSG_TYPE_HI {
		// add new user
		newUser := user.User{msgMap["From"].(string), remoteIp.String(), ""}
		user.UserManager.AddUser(&newUser)
		event.Trigger("view:msg", msg, nil)
		// response
	} else {
		fmt.Println("unknown msg: ", msgMap)
	}
}

func (this *CommunicationServer) broadcastMe() {
	this.sendMessage(net.IPv4(255, 255, 255, 255), define.Message{define.MSG_TYPE_HI, "anzhihun", "all", "hi"})
}

func (this *CommunicationServer) sendMessage(remoteIp net.IP, msg interface{}) error {
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
