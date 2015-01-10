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
		Users:   []*user.User{user.FindUser(params["creator"].(string))},
		Creator: params["creator"].(string),
	}
	user.AddChannel(&newChannel)

	// ack
	existUser := user.FindUser(newChannel.Creator)
	var ackUser = entity.WSAckUser{
		Name:     existUser.Name,
		IconUrl:  existUser.HeadImg,
		ServerId: 0,
		Online:   true,
	}

	ackChannel := entity.WSAckChannel{
		Name:    newChannel.Name,
		Creator: newChannel.Creator,
		Users:   []entity.WSAckUser{ackUser},
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
		conn.SendMsg2Client(newChannel.Creator, content)
	}

}

func (this *ChannelContext) getChannels(rw web.ResponseWriter, req *web.Request) {
	req.ParseForm()
	userName := req.Form.Get("user")
	channels := user.AllChannels()
	ackChannels := []entity.WSAckChannel{}
	for _, channel := range channels {
		if channel.ContainsUser(userName) {
			ackChannels = append(ackChannels, convert2AckChannel(channel))
		}
	}

	if content, err := json.Marshal(ackChannels); err != nil {
		http.Error(rw, "marshal ack channels error! "+err.Error(), 500)
	} else {
		rw.Write(content)
	}

}

func convert2AckChannel(channel *user.Channel) entity.WSAckChannel {
	ackUsers := []entity.WSAckUser{}
	for _, user := range channel.Users {
		ackUsers = append(ackUsers, entity.WSAckUser{
			Name:     user.Name,
			IconUrl:  user.HeadImg,
			Online:   user.Online,
			ServerId: 0,
		})
	}

	return entity.WSAckChannel{
		Name:     channel.Name,
		Creator:  channel.Creator,
		Users:    ackUsers,
		ServerId: 0,
	}
}
