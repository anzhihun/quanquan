package storage

func StoreChannels(content []byte) error {
	// store in .channels file
	return store(CHANNEL_FILE_PATH, content)
}

func ReadChannels() ([]byte, error) {
	// read from .channels file
	return read(CHANNEL_FILE_PATH)
}
