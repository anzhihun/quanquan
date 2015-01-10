package controller

import (
	"conn"
	"encoding/json"
	"entity"
	"github.com/gocraft/web"
	"io/ioutil"
	"net/http"
	"user"
	"utils"
)

func (this *ChannelContext) addChannel(rw web.ResponseWriter, req *web.Request) {
	result, _ := ioutil.ReadAll(req.Body)
	req.Body.Close()
	params, _ := utils.DecodeJsonMsg(string(result))
	newChannel := user.Channel{Name: params["name"].(string),
		Users:   []*user.User{},
		Creator: params["creator"].(string),
	}
	user.AddChannel(&newChannel)

	// ack
	ackChannel := entity.WSAckChannel{
		Name:    newChannel.Name,
		Creator: newChannel.Creator,
		Users:   []entity.WSAckUser{},
	}
	if content, err := json.Marshal(ackChannel); err != nil {
		http.Error(rw, "marshal ack channel error!", 500)
	} else {
		rw.Write(content)
	}

	//broadcast to client
	msg := entity.WSAckNewChannel{
		MsgType: entity.WS_MSGTYPE_NEW_CHANNEL,
		Channel: ackChannel,
	}

	if content, err := json.Marshal(msg); err != nil {
		http.Error(rw, "marshal broadcast channel msg error!", 500)
	} else {
		conn.Broadcast2AllClient(content)
	}

}
