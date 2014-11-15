package user

import (
	"errors"
	osuser "os/user"
	"utils"
)

type User struct {
	Name    string
	IP      string
	HeadImg string
	Online  bool
}

var osUser, _ = osuser.Current()
var ip, _ = utils.ExternalIP()
var Self = User{osUser.Username, ip, "/images/anzhihun.png", true}

func SignUp(userName, password string) error {
	if len(userName) == 0 {
		return errors.New("The user name can not be empty")
	}

	if len(password) == 0 {
		return errors.New("The password can not be empty")
	}

	if UserManager.FindUser(userName) != nil {
		return errors.New("The user is exist")
	}

	return nil
}
