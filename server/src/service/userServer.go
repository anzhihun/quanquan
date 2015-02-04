package service

import (
	"adapter"
	"entity"
	"errors"
	osuser "os/user"
	"utils"
)

var osUser, _ = osuser.Current()
var ip, _ = utils.ExternalIP()
var Self *entity.User = nil

var allUsers = []*entity.User{}

func AddUser(user *entity.User) error {

	if err := check(user.Name, user.Password); err != nil {
		return err
	}

	allUsers = append(allUsers, user)

	// save to disk
	if err := adapter.AddUser(user); err != nil {
		return err
	}

	return nil
}

func check(userName, password string) error {
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

func RemoveUser(user *entity.User) {
	for index := 0; index < len(allUsers); index++ {
		if allUsers[index].Name == user.Name {
			allUsers = append(allUsers[:index], allUsers[index+1:]...)
			return
		}
	}
	return
}

func FindUser(userName string) *entity.User {
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

func AllUser() []*entity.User {
	return allUsers
}

func Clear() {
	allUsers = allUsers[:0]
}

func initUsers() {

	adapter.InitStorage()
	var err error
	allUsers, err = adapter.GetAllUsers()
	if err != nil {
		// TODO log error
		return
	}
}

func Validate(userName, password string) bool {

	existUser := FindUser(userName)
	if existUser == nil {
		return false
	}

	if password != existUser.Password {
		return false
	}
	return true
}
