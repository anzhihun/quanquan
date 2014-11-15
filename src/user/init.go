package user

import (
	"bytes"
	"encoding/gob"
	"event"
	"storage"
)

func Init() {
	initUsers()
	event.On("user:add", storeAllUsers)
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
