package adapter

import (
	"encoding/json"
	"entity"
	"fmt"
	"github.com/fatih/structs"
	"storage"
)

func GetAllChannels() ([]*entity.Channel, error) {
	channels := []*entity.Channel{}
	channelDocuments, err := storage.TiedotDB.GetAllDocuments(COL_NAME_CHAN)
	if err != nil {
		return channels, err
	}

	for _, val := range channelDocuments {
		var channel entity.Channel
		if err = json.Unmarshal(val, &channel); err == nil {
			channels = append(channels, &channel)
		}
	}

	return channels, nil
}

func AddChannel(channel *entity.Channel) error {
	if _, err := storage.TiedotDB.AddDocument(COL_NAME_CHAN, structs.Map(channel)); err != nil {
		fmt.Println("failed to add channel.", err.Error())
		return err
	}
	return nil
}

func ClearAllChannels() error {
	err := storage.TiedotDB.DeleteCollection(COL_NAME_CHAN)
	if err != nil {
		fmt.Println("failed to clear all channels.", err.Error())
		return err
	}

	// recreate collection
	storage.TiedotDB.AddCollection(COL_NAME_CHAN)

	return err
}
