package utils

import (
	"testing"
)

func TestDecodeJsonMsg(t *testing.T) {
	testValue := "{\"Name\": \"test\"}"
	value, err := DecodeJsonMsg(testValue)
	if err != nil || value["Name"].(string) != "test" {
		t.Fatal("failed to decode json msg")
	}

	testValue = "{\"Name\"\"test\"}"
	value, err = DecodeJsonMsg(testValue)
	if err == nil {
		t.Fatal("failed to decode wrong json msg")
	}

	testValue = "{\"Name\": \"test\", \"Value\":\"testValue\"}"
	value, err = DecodeJsonMsg(testValue)
	if err != nil || value["Value"].(string) != "testValue" {
		t.Fatal("failed to decode json msg with more attributes")
	}

}
