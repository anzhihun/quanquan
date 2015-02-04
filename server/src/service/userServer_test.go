package service

import (
	"entity"
	"event"
	"testing"
)

func TestClearallUsers(t *testing.T) {
	Clear()
	if len(allUsers) > 0 {
		t.Fatal("failed to clear allUsers!")
	}
}

func TestAddUser(t *testing.T) {
	event.RunEventDispather()
	Clear()
	defer Clear()

	err := AddUser(&entity.User{Name: "test", Password: "123456", HeadImg: "/images/test.png"})
	if err != nil || allUsers == nil || len(allUsers) != 1 {
		t.Fatal("falied to add new user")
	}

	if allUsers[0].Name != "test" {
		t.Fatal("added user is wrong")
	}

	err = AddUser(&entity.User{Name: "test1", Password: "123456", HeadImg: "/images/test1.png"})
	if err != nil || len(allUsers) != 2 && allUsers[1].Name != "test1" {
		t.Fatal("failed to add more allUsers")
	}
}

func TestAddDuplicateUser(t *testing.T) {
	Clear()
	defer Clear()

	AddUser(&entity.User{Name: "test", Password: "123456", HeadImg: "/images/test.png", Online: false})
	AddUser(&entity.User{Name: "test", Password: "123456", HeadImg: "/images/test1.png", Online: true})
	if len(allUsers) != 1 {
		t.Fatal("failed to update user when adding duplicate user")
	}
}

func TestRemoveUser(t *testing.T) {
	defer Clear()
	Clear()
	AddUser(&entity.User{Name: "test", Password: "123456", HeadImg: "/images/test.png"})
	RemoveUser(&entity.User{Name: "test", HeadImg: "/images/test.png"})
	if len(allUsers) != 0 {
		t.Fatal("failed to remove user")
	}
	AddUser(&entity.User{Name: "test", Password: "123456", HeadImg: "/images/test.png"})
	AddUser(&entity.User{Name: "test1", Password: "123456", HeadImg: "/images/test2.png"})
	RemoveUser(&entity.User{Name: "test", HeadImg: "/images/test.png"})
	if len(allUsers) != 1 || allUsers[0].Name != "test1" {
		t.Fatal("failed to remove correct user")
	}
}

func TestFindUser(t *testing.T) {
	defer Clear()
	Clear()
	AddUser(&entity.User{Name: "test", Password: "123456", HeadImg: "/images/test.png"})
	findUser := FindUser("userName")
	if findUser != nil {
		t.Fatal("failed to find user")
	}

	findUser = FindUser("test")
	if findUser == nil || findUser.Name != "test" {
		t.Fatal("failed to find user")
	}
}
