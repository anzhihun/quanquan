package user

import (
	"testing"
)

func TestClearUsers(t *testing.T) {
	UserManager.Clear()
	if UserManager.users != nil && len(UserManager.users) > 0 {
		t.Fatal("failed to clear users!")
	}
}

func TestAddUser(t *testing.T) {
	UserManager.Clear()
	defer UserManager.Clear()

	UserManager.AddUser(&User{Name: "test", IP: "127.0.0.1", HeadImg: "/images/test.png"})
	if UserManager.users == nil || len(UserManager.users) != 1 {
		t.Fatal("falied to add new user")
	}

	if UserManager.users[0].Name != "test" {
		t.Fatal("added user is wrong")
	}

	UserManager.AddUser(&User{Name: "test1", IP: "127.0.0.2", HeadImg: "/images/test1.png"})
	if len(UserManager.users) != 2 && UserManager.users[1].Name != "test1" {
		t.Fatal("failed to add more users")
	}
}

func TestAddDuplicateUser(t *testing.T) {
	UserManager.Clear()
	defer UserManager.Clear()

	UserManager.AddUser(&User{Name: "test", IP: "127.0.0.1", HeadImg: "/images/test.png", Online: false})
	UserManager.AddUser(&User{Name: "test", IP: "127.0.0.1", HeadImg: "/images/test1.png", Online: true})
	if len(UserManager.users) != 1 || UserManager.users[0].Online != true {
		t.Fatal("failed to update user when adding duplicate user")
	}
}

func TestRemoveUser(t *testing.T) {
	defer UserManager.Clear()
	UserManager.Clear()
	UserManager.AddUser(&User{Name: "test", IP: "127.0.0.1", HeadImg: "/images/test.png"})
	UserManager.RemoveUser(&User{Name: "test", IP: "127.0.0.1", HeadImg: "/images/test.png"})
	if len(UserManager.users) != 0 {
		t.Fatal("failed to remove user")
	}
	UserManager.AddUser(&User{Name: "test", IP: "127.0.0.1", HeadImg: "/images/test.png"})
	UserManager.AddUser(&User{Name: "test1", IP: "127.0.0.2", HeadImg: "/images/test2.png"})
	UserManager.RemoveUser(&User{Name: "test", IP: "127.0.0.1", HeadImg: "/images/test.png"})
	if len(UserManager.users) != 1 && UserManager.users[0].Name != "test1" {
		t.Fatal("failed to remove correct user")
	}
}

func TestFindUser(t *testing.T) {
	defer UserManager.Clear()
	UserManager.Clear()
	UserManager.AddUser(&User{Name: "test", IP: "127.0.0.1", HeadImg: "/images/test.png"})
	findUser := UserManager.FindUser("userName")
	if findUser != nil {
		t.Fatal("failed to find user")
	}

	findUser = UserManager.FindUser("test")
	if findUser == nil || findUser.Name != "test" {
		t.Fatal("failed to find user")
	}
}
