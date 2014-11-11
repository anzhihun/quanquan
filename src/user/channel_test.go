package user

import (
	"testing"
)

func TestAddChannelUser(t *testing.T) {
	channel := Channel{name: "test", users: nil}
	channel.AddUser(&User{Name: "name1", IP: "127.0.0.1", HeadImg: "/images/test1.png"})

	if channel.users == nil || len(channel.users) != 1 || channel.users[0].Name != "name1" {
		t.Fatal("failed to add channel user")
	}
}

func TestRemoveChannelUser(t *testing.T) {
	channel := Channel{name: "test", users: nil}
	channel.AddUser(&User{Name: "name1", IP: "127.0.0.1", HeadImg: "/images/test1.png"})
	channel.RemoveUser(&User{Name: "name1", IP: "127.0.0.1", HeadImg: "/images/test1.png"})
	if len(channel.users) != 0 {
		t.Fatal("failed to remove channel user")
	}
}

func TestAddChannel(t *testing.T) {
	channels = channels[:0]
	AddChannel(&Channel{name: "test", users: nil})
	if len(channels) != 1 || channels[0].name != "test" {
		t.Fatal("failed to add new channel")
	}

	AddChannel(&Channel{name: "test1", users: nil})
	if len(channels) != 2 || channels[1].name != "test1" {
		t.Fatal("failed to add more channel")
	}
}

func TestRemoveChannel(t *testing.T) {
	channels = channels[:0]
	AddChannel(&Channel{name: "test", users: nil})
	RemoveChannelByName("test")
	if len(channels) != 0 {
		t.Fatal("failed to remove channel")
	}

	AddChannel(&Channel{name: "test", users: nil})
	AddChannel(&Channel{name: "test1", users: nil})
	RemoveChannelByName("test")
	if len(channels) != 1 || channels[0].name != "test1" {
		t.Fatal("failed to remove one channel of multiple")
	}
}

func TestFindChannelByName(t *testing.T) {
	channels = channels[:0]
	channel := FindChannelByName("test")
	if channel != nil {
		t.Fatal("find not exist channel")
	}
	AddChannel(&Channel{name: "test", users: nil})
	AddChannel(&Channel{name: "test1", users: nil})

	channel = FindChannelByName("test")
	if channel == nil || channel.name != "test" {
		t.Fatal("failed to find exist channel")
	}
}
