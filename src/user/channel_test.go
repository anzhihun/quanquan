package user

import (
	"testing"
)

func TestAddChannelUser(t *testing.T) {
	channel := Channel{Name: "test", Users: nil}
	channel.AddUser(&User{Name: "name1", IP: "127.0.0.1", HeadImg: "/images/test1.png"})

	if channel.Users == nil || len(channel.Users) != 1 || channel.Users[0].Name != "name1" {
		t.Fatal("failed to add channel user")
	}
}

func TestRemoveChannelUser(t *testing.T) {
	channel := Channel{Name: "test", Users: nil}
	channel.AddUser(&User{Name: "name1", IP: "127.0.0.1", HeadImg: "/images/test1.png"})
	channel.RemoveUser(&User{Name: "name1", IP: "127.0.0.1", HeadImg: "/images/test1.png"})
	if len(channel.Users) != 0 {
		t.Fatal("failed to remove channel user")
	}
}

func TestAddChannel(t *testing.T) {
	channels = channels[:0]
	AddChannel(&Channel{Name: "test", Users: nil})
	if len(channels) != 1 || channels[0].Name != "test" {
		t.Fatal("failed to add new channel")
	}

	AddChannel(&Channel{Name: "test1", Users: nil})
	if len(channels) != 2 || channels[1].Name != "test1" {
		t.Fatal("failed to add more channel")
	}
}

func TestRemoveChannel(t *testing.T) {
	channels = channels[:0]
	AddChannel(&Channel{Name: "test", Users: nil})
	RemoveChannelByName("test")
	if len(channels) != 0 {
		t.Fatal("failed to remove channel")
	}

	AddChannel(&Channel{Name: "test", Users: nil})
	AddChannel(&Channel{Name: "test1", Users: nil})
	RemoveChannelByName("test")
	if len(channels) != 1 || channels[0].Name != "test1" {
		t.Fatal("failed to remove one channel of multiple")
	}
}

func TestFindChannelByName(t *testing.T) {
	channels = channels[:0]
	channel := FindChannelByName("test")
	if channel != nil {
		t.Fatal("find not exist channel")
	}
	AddChannel(&Channel{Name: "test", Users: nil})
	AddChannel(&Channel{Name: "test1", Users: nil})

	channel = FindChannelByName("test")
	if channel == nil || channel.Name != "test" {
		t.Fatal("failed to find exist channel")
	}
}
