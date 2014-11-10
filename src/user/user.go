package user

import (
	osuser "os/user"
	"utils"
)

type User struct {
	Name    string
	IP      string
	HeadImg string
}

var osUser, _ = osuser.Current()
var ip, _ = utils.ExternalIP()
var Self User = User{osUser.Username, ip, "/images/anzhihun.png"}
