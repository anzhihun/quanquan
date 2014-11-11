package storage

import (
	"utils"
)

func StoreChannels(content []byte) error {
	// store in .channels file
	if !utils.IsDirExist(CONFIG_DIR) {
		if err := utils.CreateDir(CONFIG_DIR); err != nil {
			return err
		}
	}

	if !utils.IsFileExist(CHANNEL_FILE_PATH) {
		if err := utils.CreateFile(CHANNEL_FILE_PATH); err != nil {
			return err
		}
	}

	return utils.WriteFile(CHANNEL_FILE_PATH, content)
}

func ReadChannels() ([]byte, error) {
	// read from .channels file
	if !utils.IsDirExist(CONFIG_DIR) || !utils.IsFileExist(CHANNEL_FILE_PATH) {
		return []byte{}, nil
	}

	return utils.ReadFile(CHANNEL_FILE_PATH)
}
