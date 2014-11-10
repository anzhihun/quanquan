package user

import (
	osuser "os/user"
)

type User struct {
	Name    string
	IP      string
	HeadImg string
}

var osUser, _ = osuser.Current()
var Self User = User{osUser.Username, "127.0.0.1", "/images/anzhihun.png"}
