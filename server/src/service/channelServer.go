package service

import (
	"bytes"
	"encoding/gob"
	"entity"
	"storage"
)

var channels = make([]*entity.Channel, 0)

func AddChannel(channel *entity.Channel) {
	channels = append(channels, channel)
	storeAllChannel(nil, nil)
}

func RemoveChannelByName(chanName string) {
	for i := 0; i < len(channels); i++ {
		if channels[i].Name == chanName {
			channels = append(channels[:i], channels[i+1:]...)
			break
		}
	}
}

func FindChannelByName(chanName string) *entity.Channel {
	for i := 0; i < len(channels); i++ {
		if channels[i].Name == chanName {
			return channels[i]
		}
	}
	return nil
}

func AllChannels() []*entity.Channel {
	return channels
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
