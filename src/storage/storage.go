package storage

import (
	"path/filepath"
)

import (
	"utils"
)

const (
	CONFIG_DIR = ".conf"

	CHANNEL_FILE_NAME = ".channels"
	CHANNEL_FILE_PATH = CONFIG_DIR + string(filepath.Separator) + CHANNEL_FILE_NAME

	USER_FILE_NAME = ".users"
	USER_FILE_PATH = CONFIG_DIR + string(filepath.Separator) + USER_FILE_NAME
)

func store(configFilePath string, content []byte) error {
	if !utils.IsDirExist(CONFIG_DIR) {
		if err := utils.CreateDir(CONFIG_DIR); err != nil {
			return err
		}
	}

	if !utils.IsFileExist(configFilePath) {
		if err := utils.CreateFile(configFilePath); err != nil {
			return err
		}
	}

	return utils.WriteFile(configFilePath, content)
}

func read(configFilePath string) ([]byte, error) {
	if !utils.IsDirExist(CONFIG_DIR) || !utils.IsFileExist(configFilePath) {
		return []byte{}, nil
	}

	return utils.ReadFile(configFilePath)
}
