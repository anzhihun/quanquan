package service

import (
	"define"
	"encoding/json"
	"event"
	// "fmt"
	"net"
	"testing"
	"time"
	"user"
)

func TestHandleOnlineMessage(t *testing.T) {

	event.RunEventDispather()
	user.UserManager.Clear()
	defer user.UserManager.Clear()

	channel := make(chan interface{})
	eventId := event.On("view:msg", func(newValue, oldValue interface{}) {
		defer func() { channel <- 1 }()
		msg := string(newValue.([]byte))
		if msg != "{\"MsgType\":\"online\",\"From\":\"testName1\",\"HeadImg\":\"testImg1\",\"To\":\"all\",\"IsPublic\":true,\"Content\":\"testContent1\"}" {
			t.Fatal("Handle online msg error!")
		}
		newUser := user.UserManager.FindUser("testName1")
		if newUser == nil || newUser.IP != "255.255.255.255" {
			t.Fatal("Handle online message not add new user")
		}
	})
	defer event.Off(eventId)

	msg := define.Message{
		MsgType:  define.MSG_TYPE_ONLINE,
		From:     "testName1",
		HeadImg:  "testImg1",
		To:       "all",
		IsPublic: true,
		Content:  "testContent1",
	}

	msgBytes, _ := json.Marshal(msg)
	commServer.handleMessage(net.IPv4(255, 255, 255, 255), msgBytes)
	select {
	case <-channel:
	case <-time.After(10 * time.Second):
		t.Fatal("msg not callback")
	}
}

func TestHandleOfflineMessage(t *testing.T) {
	event.RunEventDispather()
	user.UserManager.Clear()
	defer user.UserManager.Clear()

	msgCount := 0
	channel := make(chan interface{})
	eventId := event.On("view:msg", func(newValue, oldValue interface{}) {
		defer func() {
			msgCount = msgCount + 1
			if msgCount == 2 {
				channel <- 1
			}
		}()
		msg := string(newValue.([]byte))
		if msgCount == 1 && msg != "{\"MsgType\":\"offline\",\"From\":\"testName1\",\"HeadImg\":\"testImg1\",\"To\":\"all\",\"IsPublic\":true,\"Content\":\"testContent1\"}" {
			t.Fatal("Handle offline msg error!")
		}
		newUser := user.UserManager.FindUser("testName1")
		if msgCount == 1 && newUser != nil {
			t.Fatal("Handle offline message not remove user")
		}
	})
	defer event.Off(eventId)

	msg := define.Message{
		MsgType:  define.MSG_TYPE_ONLINE,
		From:     "testName1",
		HeadImg:  "testImg1",
		To:       "all",
		IsPublic: true,
		Content:  "testContent1",
	}

	msgBytes, _ := json.Marshal(msg)
	commServer.handleMessage(net.IPv4(255, 255, 255, 255), msgBytes)

	msg.MsgType = define.MSG_TYPE_OFFLINE
	msgBytes, _ = json.Marshal(msg)
	commServer.handleMessage(net.IPv4(255, 255, 255, 255), msgBytes)

	select {
	case <-channel:
	case <-time.After(10 * time.Second):
		t.Fatal("msg not callback")
	}
}

func TestHandleJoinMessage(t *testing.T) {
	event.RunEventDispather()

	channel := make(chan interface{})
	eventId := event.On("view:msg", func(newValue, oldValue interface{}) {
		defer func() { channel <- 1 }()
		msg := string(newValue.([]byte))
		if msg != "{\"MsgType\":\"join\",\"From\":\"testName1\",\"HeadImg\":\"testImg1\",\"To\":\"all\",\"IsPublic\":true,\"Content\":\"global\"}" {
			t.Fatal("Handle join msg error!")
		}
	})
	defer event.Off(eventId)

	msg := define.Message{
		MsgType:  define.MSG_TYPE_JOIN,
		From:     "testName1",
		HeadImg:  "testImg1",
		To:       "all",
		IsPublic: true,
		Content:  "global",
	}

	msgBytes, _ := json.Marshal(msg)
	commServer.handleMessage(net.IPv4(255, 255, 255, 255), msgBytes)
	select {
	case <-channel:
	case <-time.After(10 * time.Second):
		t.Fatal("msg not callback")
	}
}

func TestHandleTalkMessage(t *testing.T) {
	event.RunEventDispather()

	channel := make(chan interface{})
	eventId := event.On("view:msg", func(newValue, oldValue interface{}) {
		defer func() { channel <- 1 }()
		msg := string(newValue.([]byte))
		if msg != "{\"MsgType\":\"talk\",\"From\":\"testName1\",\"HeadImg\":\"testImg1\",\"To\":\"all\",\"IsPublic\":true,\"Content\":\"global\"}" {
			t.Fatal("Handle talk msg error!")
		}
	})
	defer event.Off(eventId)

	msg := define.Message{
		MsgType:  define.MSG_TYPE_TALK,
		From:     "testName1",
		HeadImg:  "testImg1",
		To:       "all",
		IsPublic: true,
		Content:  "global",
	}

	msgBytes, _ := json.Marshal(msg)
	commServer.handleMessage(net.IPv4(255, 255, 255, 255), msgBytes)
	select {
	case <-channel:
	case <-time.After(10 * time.Second):
		t.Fatal("msg not callback")
	}
}
