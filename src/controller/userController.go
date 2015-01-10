package controller

import (
	// "define"
	"encoding/json"
	"entity"
	"fmt"
	// "event"
	"conn"
	"github.com/gocraft/web"
	"io/ioutil"
	"net/http"
	"user"
	"utils"
)

func (this *UserContext) Login(rw web.ResponseWriter, req *web.Request) {
	result, _ := ioutil.ReadAll(req.Body)
	req.Body.Close()
	params, err := utils.DecodeJsonMsg(string(result))
	if err != nil {
		http.Error(rw, err.Error(), 404)
		return
	}

	if !user.Validate(params["name"].(string), params["password"].(string)) {
		http.Error(rw, "invalid user name or password", 500)
	}

	var ackContent = []byte{}
	if ackContent, err = this.getAckHttpUser(params["name"].(string)); err != nil {
		http.Error(rw, "marshal user error!", 500)
	} else {
		rw.Write(ackContent)
	}

	this.sendOnlineMsg(params["name"].(string))

	// send msg to back end service
	//event.Trigger(event.EVENT_F2B_LOGIN, params["name"].(string), nil)

}

func (this *UserContext) getAckHttpUser(userName string) (content []byte, err error) {
	existUser := user.FindUser(userName)

	var ackUser = entity.WSAckUser{
		Name:     existUser.Name,
		IconUrl:  existUser.HeadImg,
		ServerId: 0,
		Online:   true,
	}

	content, err = json.Marshal(ackUser)
	return
}

func (this *UserContext) sendOnlineMsg(userName string) {
	msg := entity.WSAckTalk{
		MsgType:     entity.WS_MSGTYPE_ONLINE,
		ContentType: entity.WS_MSGCONTENT_TYPE_TEXT,
		Sender:      "system",
		Is2P:        false,
		Receiver:    "all",
		Content:     userName + " come in",
	}

	if content, err := json.Marshal(msg); err != nil {
		fmt.Println("send online msg error! ", err.Error())
		return
	} else {
		conn.Broadcast2AllClient(content)
	}

}

func (this *UserContext) SignUp(rw web.ResponseWriter, req *web.Request) {
	result, _ := ioutil.ReadAll(req.Body)
	req.Body.Close()
	params, err := utils.DecodeJsonMsg(string(result))
	if err != nil {
		http.Error(rw, err.Error(), 404)
		return
	}

	userName := params["name"].(string)
	password := params["password"].(string)
	if err = user.SignUp(userName, password); err != nil {
		http.Error(rw, err.Error(), 500)
		return
	}

	user.AddUser(&user.User{
		Name:     userName,
		Password: password,
		HeadImg:  "/images/defaultHead.png",
		Online:   true,
	})

	var ackContent = []byte{}
	if ackContent, err = this.getAckHttpUser(userName); err != nil {
		http.Error(rw, "marshal user error!", 500)
	} else {
		rw.Write(ackContent)
	}

	ackNewUser := entity.WSAckNewUser{
		MsgType: entity.WS_MSGTYPE_NEW_USER,
		User: entity.WSAckUser{
			Name:     userName,
			IconUrl:  "/images/defaultHead.png",
			ServerId: 0,
			Online:   true,
		},
	}
	if ackContent, err = json.Marshal(ackNewUser); err != nil {
		fmt.Println("marshal new user msg error!", err.Error())
	} else {
		conn.Broadcast2AllClient(ackContent)
	}

	this.sendOnlineMsg(userName)
	// send msg to back end service
	//event.Trigger(event.EVENT_F2B_ADD_USER, define.AddUserMessage{UserName: userName, Password: password}, nil)

	// add new user
	// ip, _ := utils.ExternalIP()
	// user.AddUser(&user.User{
	// 	Name:     userName,
	// 	Password: password,
	// 	HeadImg:  "/images/defaultHead.png",
	// 	IP:       ip,
	// 	Online:   true,
	// })
}

func (this *UserContext) GetUsers(rw web.ResponseWriter, req *web.Request) {
	users := user.AllUser()
	ackUser := []entity.WSAckUser{}
	for _, user := range users {
		ackUser = append(ackUser, entity.WSAckUser{
			Name:     user.Name,
			IconUrl:  user.HeadImg,
			Online:   user.Online,
			ServerId: 0,
		})
	}

	contents, err := json.Marshal(ackUser)
	if err != nil {
		rw.Write([]byte("error: " + err.Error()))
	} else {
		rw.Write(contents)
	}
	// req.ParseForm()
}
