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

	if err = user.SignUp(params["name"].(string), params["password"].(string)); err != nil {
		http.Error(rw, err.Error(), 500)
	}
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
