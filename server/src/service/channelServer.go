package service

import (
	"adapter"
	"entity"
	"fmt"
)

var channels = make([]*entity.Channel, 0)

func AddChannel(channel *entity.Channel) {
	channels = append(channels, channel)
	if err := adapter.AddChannel(channel); err != nil {
		fmt.Println("failed to store channel", err.Error())
	}
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

	var err error
	channels, err = adapter.GetAllChannels()
	if err != nil {
		// TODO log error
		return
	}
}
