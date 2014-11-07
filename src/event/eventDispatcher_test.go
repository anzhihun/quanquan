package event

import (
	"strings"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	RunEventDispather()
	startListenerMonitor()
	channel := make(chan interface{})
	eventId := On("msgType", func(newValue, oldValue interface{}) {
		defer func() {
			channel <- 12
		}()

		if !strings.EqualFold(newValue.(string), "newValue") {
			t.Fatal("callback new value is not correct!")
		}
		if !strings.EqualFold(oldValue.(string), "oldValue") {
			t.Fatal("callback old value is not correct!")
		}
	})
	defer Off(eventId)

	Trigger("msgType", "newValue", "oldValue")
	select {
	case <-channel:
	case <-time.After(10 * time.Second):
		t.Fatal("msg not callback")
	}
}
