package adapter

import (
	"entity"
	"testing"
)

func TestGetAllChannels(t *testing.T) {
	InitStorage()
	ClearAllChannels()
	channels, err := GetAllChannels()
	if err != nil {
		t.Fatal("failed to get all channel.", err.Error())
	}
	if len(channels) != 0 {
		t.Fatal("the channels should be empty on inital.")
	}

	channel := entity.Channel{Name: "test", Users: nil}
	user := entity.User{
		Name:     "test",
		Password: "123456",
		HeadImg:  "/images/test.png",
	}
	channel.AddUser(&user)

	err = AddChannel(&channel)
	channels, err = GetAllChannels()
	if err != nil {
		t.Fatal("failed to get all channel again.", err.Error())
	}
	if len(channels) != 1 {
		t.Fatal("the channels should contains one.")
	}
	if len(channels[0].Users) != 1 {
		t.Fatal("the channel should be contains one user")
	}
	if channels[0].Users[0].Name != "test" || channels[0].Users[0].Password != "123456" || channels[0].Users[0].HeadImg != "/images/test.png" {
		t.Fatal("the channel's user is wrong")
	}

}
