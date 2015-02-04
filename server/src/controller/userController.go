package controller

import (
	"conn"
	"define"
	"encoding/json"
	"entity"
	"fmt"
	"github.com/gocraft/web"
	"i18n"
	"io/ioutil"
	"net/http"
	"service"
	"user"
	"utils"
)

type UserContext struct {
	*RootContext
}

func (this *UserContext) Login(rw web.ResponseWriter, req *web.Request) {
	result, _ := ioutil.ReadAll(req.Body)
	req.Body.Close()
	params, err := utils.DecodeJsonMsg(string(result))
	if err != nil {
		http.Error(rw, err.Error(), 404)
		return
	}

	if !service.Validate(params["name"].(string), params["password"].(string)) {
		http.Error(rw, "invalid user name or password", 500)
	}

	var ackContent = []byte{}
	if ackContent, err = getAckHttpUser(params["name"].(string)); err != nil {
		http.Error(rw, "marshal user error!", 500)
	} else {
		rw.Write(ackContent)
	}

	sendOnlineMsg(params["name"].(string))

	// send msg to back end service
	//event.Trigger(event.EVENT_F2B_LOGIN, params["name"].(string), nil)

}

func getAckHttpUser(userName string) (content []byte, err error) {
	existUser := service.FindUser(userName)

	var ackUser = entity.WSAckUser{
		Name:     existUser.Name,
		IconUrl:  existUser.HeadImg,
		ServerId: 0,
		Online:   true,
	}

	content, err = json.Marshal(ackUser)
	return
}

func sendOnlineMsg(userName string) {
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
	langId := utils.GetLanguageId(req.Header)
	translator := i18n.GetTranslator(langId)

	userName, password, errInfo, errCode := prepareSignUpParams(req, translator)
	if errCode != 0 {
		http.Error(rw, errInfo, errCode)
		return
	}

	if errInfo, errCode = addNewUser(userName, password, translator); errCode != 0 {
		http.Error(rw, errInfo, errCode)
		return
	}

	responseSignUp(rw, userName, translator)
	broadcastSignUpMsg(userName)
}

func prepareSignUpParams(req *web.Request, translator *i18n.Translator) (userName string, password string, errInfo string, errCode int) {
	errCode = 0
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		errInfo = translator.T("ERR_READ_HTTP_BODY", map[string]interface{}{
			"Error": err.Error(),
		})
		errCode = define.ERR_READ_HTTP_BODY
		return
	}
	req.Body.Close()

	var params map[string]interface{}
	if params, err = utils.DecodeJsonMsg(string(result)); err != nil {
		errInfo = translator.T("ERR_DECODE_JSON", map[string]interface{}{
			"Error": err.Error(),
		})
		errCode = define.ERR_DECODE_JSON
		return
	}

	var ok bool
	if userName, ok = params["name"].(string); !ok {
		errInfo = translator.T("ERR_WRONG_TYPE", map[string]interface{}{
			"Var":  "name",
			"Type": "string",
		})
		errCode = define.ERR_WRONG_TYPE
		return
	}

	if password, ok = params["password"].(string); !ok {
		errInfo = translator.T("ERR_WRONG_TYPE", map[string]interface{}{
			"Var":  "password",
			"Type": "string",
		})
		errCode = define.ERR_WRONG_TYPE
		return
	}

	return
}

func addNewUser(userName string, password string, translator *i18n.Translator) (errInfo string, errCode int) {
	errCode = 0
	if err := service.AddUser(&entity.User{
		Name:     userName,
		Password: password,
		HeadImg:  "/images/defaultHead.png",
		Online:   true,
	}); err != nil {
		errInfo = translator.T("ERR_ADD_USER", map[string]interface{}{
			"Error": err.Error(),
		})
		errCode = define.ERR_INVALID_USER
	}

	return
}

func responseSignUp(rw web.ResponseWriter, userName string, translator *i18n.Translator) {
	if ackContent, err := getAckHttpUser(userName); err != nil {
		errInfo := translator.T("ERR_ENCODE_JSON", map[string]interface{}{
			"Error": err.Error(),
		})
		http.Error(rw, errInfo, define.ERR_ENCODE_JSON)
		return
	} else {
		rw.Write(ackContent)
	}
}

func broadcastSignUpMsg(userName string) {
	newUserMsg := entity.WSAckNewUser{
		MsgType: entity.WS_MSGTYPE_NEW_USER,
		User: entity.WSAckUser{
			Name:     userName,
			IconUrl:  "/images/defaultHead.png",
			ServerId: 0,
			Online:   true,
		},
	}
	if ackContent, err := json.Marshal(newUserMsg); err != nil {
		fmt.Println("marshal new user msg error!", err.Error())
	} else {
		conn.Broadcast2AllClient(ackContent)
	}

	sendOnlineMsg(userName)
}

func (this *UserContext) GetUsers(rw web.ResponseWriter, req *web.Request) {
	req.ParseForm()
	channelName := req.Form.Get("channel")
	fmt.Println("get user's channel = ", channelName)
	var users []*entity.User
	if channelName == "Global" {
		users = service.AllUser()
	} else {
		users = user.FindChannelByName(channelName).Users
	}

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
