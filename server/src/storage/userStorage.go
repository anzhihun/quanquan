package storage

func StoreUsers(content []byte) error {
	return store(USER_FILE_PATH, content)
}

func ReadUsers() ([]byte, error) {
	return read(USER_FILE_PATH)
}
