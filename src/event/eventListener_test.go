package event

import (
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestIsEventsFull(t *testing.T) {
	var locker sync.Mutex
	listener := eventListener{eventType: "msgType", id: 1, eventFunc: nil, cond: sync.NewCond(&locker)}

	for i := 0; i < 1001; i++ {
		listener.onEvent(eventMsg{"msgType", strconv.FormatInt(int64(i), 10), strconv.FormatInt(int64(i), 10)})
	}

	if !listener.isEventsFull() {
		t.Fatal("it should be full after push 1001 event to it.")
	}
}

func TestOnEventAndListen(t *testing.T) {
	startListenerMonitor()
	msgChan := make(chan int, 1001)
	index := 0

	testFunc := func(oldValue, newValue interface{}) {
		msgChan <- index
		index = index + 1
	}

	var locker sync.Mutex
	listener := eventListener{eventType: "msgType", id: 1, eventFunc: testFunc, cond: sync.NewCond(&locker)}
	listener.listen()

	for i := 0; i < 999; i++ {
		listener.onEvent(eventMsg{"msgType", strconv.FormatInt(int64(i), 10), strconv.FormatInt(int64(i), 10)})
	}

	var result int
	for i := 0; i < 999; i++ {
		result = <-msgChan
		if result != i {
			t.Fatal("listener shoule be called on receiving event and execute orderly!")
		}
	}

	listener.StopProcess()
	time.Sleep(100 * time.Millisecond)
	if listener.stop {
		t.Fatal("listener should stop when call StopProcess")
	}
}

func TestOnEventAndListen_concurrent(t *testing.T) {
	startListenerMonitor()
	msgChan := make(chan int, 1001)
	index := 0

	testFunc := func(oldValue, newValue interface{}) {
		msgChan <- index
		index = index + 1
	}

	var locker sync.Mutex
	listener := eventListener{eventType: "msgType", id: 1, eventFunc: testFunc, cond: sync.NewCond(&locker)}
	listener.listen()

	for i := 0; i < 999; i++ {
		go listener.onEvent(eventMsg{"msgType", strconv.FormatInt(int64(i), 10), strconv.FormatInt(int64(i), 10)})
	}

	var result int
	for i := 0; i < 999; i++ {
		result = <-msgChan
		if result != i {
			t.Fatal("listener shoule be called on receiving event and execute orderly!")
		}
	}
}
