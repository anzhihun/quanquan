package user

import (
	"bytes"
	"encoding/gob"
	"event"
	"storage"
)

func Init() {
	initUsers()
	initChannels()
	event.On("user:add", storeAllUsers)
	event.On("channel:add", storeAllChannel)
	event.On("channel:user:change", storeAllChannel)
}

func initUsers() {
	userContent, err := storage.ReadUsers()
	if err != nil {
		// TODO log error
		return
	}

	userBuffer := bytes.NewBuffer(userContent)
	enc := gob.NewDecoder(userBuffer)
	if err = enc.Decode(&allUsers); err != nil {
		//TODO log error
		return
	}
}

func storeAllUsers(newValue, oldValue interface{}) {
	// save all user
	var userBuffer bytes.Buffer
	enc := gob.NewEncoder(&userBuffer)
	err := enc.Encode(allUsers)
	if err != nil {
		// TODO log error
		return
	}

	storage.StoreUsers(userBuffer.Bytes())
}

func initChannels() {
	channelContent, err := storage.ReadChannels()
	if err != nil {
		// TODO log error
		return
	}

	channelBuffer := bytes.NewBuffer(channelContent)
	enc := gob.NewDecoder(channelBuffer)
	if err = enc.Decode(&channels); err != nil {
		//TODO log error
		return
	}
}

func storeAllChannel(ewValue, oldValue interface{}) {
	// save all channels
	var channelBuffer bytes.Buffer
	enc := gob.NewEncoder(&channelBuffer)
	err := enc.Encode(channels)
	if err != nil {
		// TODO log error
		return
	}

	storage.StoreUsers(channelBuffer.Bytes())
}
