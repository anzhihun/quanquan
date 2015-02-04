package user

import (
	"bytes"
	"encoding/gob"
	"event"
	"storage"
)

func Init() {
	initChannels()
	event.On("channel:add", storeAllChannel)
	event.On("channel:user:change", storeAllChannel)
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

	storage.StoreChannels(channelBuffer.Bytes())
}
