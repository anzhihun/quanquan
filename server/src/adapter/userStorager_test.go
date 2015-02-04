package adapter

import (
	"entity"
	"testing"
)

func TestGetAllUsers(t *testing.T) {
	InitStorage()
	ClearAllUsers()

	users, err := GetAllUsers()
	if err != nil {
		t.Fatal("failed to get all users. ", err.Error())
	}

	if len(users) != 0 {
		t.Fatal("There should be no user initial")
	}

	if err = AddUser(&entity.User{
		Name:     "test",
		Password: "123456",
		HeadImg:  "/images/test.png",
	}); err != nil {
		t.Fatal("failed to add user.", err.Error())
	}

	users, err = GetAllUsers()
	if err != nil {
		t.Fatal("failed to get all user again.", err.Error())
	}

	if len(users) != 1 || users[0].Name != "test" || users[0].Password != "123456" || users[0].HeadImg != "/images/test.png" {
		t.Fatal("failed to get correct user.")
	}

	// ADD another user
	if err = AddUser(&entity.User{
		Name:     "test2",
		Password: "123457",
		HeadImg:  "/images/test2.png",
	}); err != nil {
		t.Fatal("failed to add user2.", err.Error())
	}

	users, err = GetAllUsers()
	if err != nil {
		t.Fatal("failed to get all user again .", err.Error())
	}

	if len(users) != 2 {
		t.Fatal("failed to get correct user.")
	}
}
