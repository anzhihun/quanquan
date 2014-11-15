package controller

import (
	"encoding/json"
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
	}

	// add new user
	ip, _ := utils.ExternalIP()
	user.AddUser(&user.User{
		Name:     userName,
		Password: password,
		HeadImg:  "/images/defaultHead.png",
		IP:       ip,
		Online:   true,
	})
}

func (this *UserContext) GetUsers(rw web.ResponseWriter, req *web.Request) {
	users := user.AllUser()
	contents, err := json.Marshal(users)
	if err != nil {
		rw.Write([]byte("error: " + err.Error()))
	} else {
		rw.Write(contents)
	}
	// req.ParseForm()
}
