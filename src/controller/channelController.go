package controller

import (
	"github.com/gocraft/web"
	"io/ioutil"
	"user"
	"utils"
)

func (this *ChannelContext) addChannel(rw web.ResponseWriter, req *web.Request) {
	result, _ := ioutil.ReadAll(req.Body)
	req.Body.Close()
	params, _ := utils.DecodeJsonMsg(string(result))
	newChannel := user.Channel{params["name"].(string), []*user.User{}}
	user.AddChannel(&newChannel)
}
