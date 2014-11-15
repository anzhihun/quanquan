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

var allUsers = []*User{&Self}

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
	} else {
		oldUser.HeadImg = user.HeadImg
		oldUser.IP = user.IP
		oldUser.Name = user.Name
		oldUser.Online = user.Online
	}
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

func AllUser() []*User {
	return allUsers
}

func Clear() {
	allUsers = allUsers[:0]
}
