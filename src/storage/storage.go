package storage

import (
	"path/filepath"
)

const (
	CONFIG_DIR        = ".conf"
	CHANNEL_FILE_NAME = ".channels"
	CHANNEL_FILE_PATH = CONFIG_DIR + string(filepath.Separator) + CHANNEL_FILE_NAME
)
