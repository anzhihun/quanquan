package user

import (
	"errors"
	"event"
	osuser "os/user"
	"utils"
)

type User struct {
	Name     string
	Password string
	IP       string
	HeadImg  string
	Online   bool
}

var osUser, _ = osuser.Current()
var ip, _ = utils.ExternalIP()
var Self *User = nil

var allUsers = []*User{}

func SignUp(userName, password string) error {
	if len(userName) == 0 {
		return errors.New("The user name can not be empty")
	}

	if len(password) == 0 {
		return errors.New("The password can not be empty")
	}

	if FindUser(userName) != nil {
		return errors.New("The user is exist")
	}

	return nil
}

func AddUser(user *User) {

	oldUser := FindUser(user.Name)
	if oldUser == nil {
		allUsers = append(allUsers, user)
	}

	// trigger msg
	event.Trigger("user:add", user, nil)
}

func RemoveUser(user *User) {
	for index := 0; index < len(allUsers); index++ {
		if allUsers[index].Name == user.Name {
			allUsers = append(allUsers[:index], allUsers[index+1:]...)
			return
		}
	}
	return
}

func FindUser(userName string) *User {
	for index := 0; index < len(allUsers); index++ {
		if allUsers[index].Name == userName {
			return allUsers[index]
		}
	}
	return nil
}

func Online(userName string) {
	if existUser := FindUser(userName); existUser != nil {
		existUser.Online = true
	}
}

func AllUser() []*User {
	return allUsers
}

func Clear() {
	allUsers = allUsers[:0]
}
