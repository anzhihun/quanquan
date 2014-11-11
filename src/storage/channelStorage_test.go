package storage

import (
	"os"
	"testing"
)

func TestReadChannels(t *testing.T) {
	_, err := ReadChannels()
	if err != nil {
		t.Fatal("can't read not exist channels")
	}
}

func TestStoreChannels(t *testing.T) {
	os.RemoveAll(CONFIG_DIR)
	defer func() {
		err := os.RemoveAll(CONFIG_DIR)
		if err != nil {
			t.Fatal("failed to remove all config folder", err)
		}
	}()

	err := StoreChannels([]byte("testcontent"))
	if err != nil {
		t.Fatal("can't store channels", err.Error())
	}

	var content []byte
	content, err = ReadChannels()
	if err != nil {
		t.Fatal("failed to read content of channel file", err.Error())
	}
	if string(content) != "testcontent" {
		t.Fatal("readed channel data is wrong. ", string(content), "is not equalto \"testcontent\"")
	}
}
