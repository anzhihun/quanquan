package event

import (
	// "fmt"
	"strings"
	"testing"
	"time"
)

func TestBindEventLister(t *testing.T) {
	defer manager.clearAllListener()
	manager.clearAllListener()

	if manager.getListenerCount("MsgType") != 0 {
		t.Fatal("获取监听器个数失败")
	}

	id1 := manager.bindEventLister("MsgType", func(newValue, oldValue interface{}) {
	})
	defer manager.unbindEventLister(id1)
	if manager.getListenerCount("MsgType") != 1 {
		t.Fatal("添加监听器失败!")
	}

	id2 := manager.bindEventLister("MsgType", func(newValue, oldValue interface{}) {
	})
	defer manager.unbindEventLister(id2)
	if manager.getListenerCount("MsgType") != 2 {
		t.Fatal("添加监听器失败!")
	}

	id3 := manager.bindEventLister("MsgType2", func(newValue, oldValue interface{}) {
	})
	defer manager.unbindEventLister(id3)
	if manager.getListenerCount("MsgType2") != 1 {
		t.Fatal("添加监听器失败!")
	}
	if manager.getListenerCount("MsgType") != 2 {
		t.Fatal("添加监听器失败!")
	}
}

func TestClearAllListeners(t *testing.T) {
	defer manager.clearAllListener()

	manager.bindEventLister("msgType", func(newValue, oldValue interface{}) {
	})

	manager.clearAllListener()
	if manager.getListenerCount("msgType") != 0 {
		t.Fatal("清除所有监听器不正确！")
	}
}

func TestUnbindEventListener(t *testing.T) {
	defer manager.clearAllListener()

	id := manager.bindEventLister("msgType", func(newValue, oldValue interface{}) {
	})
	manager.unbindEventLister(id)
	if manager.getListenerCount("msgType") != 0 {
		t.Fatal("注销监听器不正确！")
	}
}

func TestTriggerEventListener(t *testing.T) {
	defer manager.clearAllListener()

	startListenerMonitor()
	syncChan := make(chan interface{})

	manager.bindEventLister("msgType", func(newValue, oldValue interface{}) {
		if !strings.EqualFold(newValue.(string), "newValue") {
			t.Fatal("回调参数newvalue错误")
		}
		if !strings.EqualFold(oldValue.(string), "oldValue") {
			t.Fatal("回调参数oldValue错误")
		}
		syncChan <- 1
	})

	// send test msg
	manager.triggerEventListener(eventMsg{"msgType", "newValue", "oldValue"})

	select {
	case <-syncChan:
		break
	case <-time.After(10 * time.Second):
		t.Fatal("监听消息处理超时")
		break
	}

}

// 测试监听器阻塞后续监听器的情况
func TestTriggerMoreEventListener(t *testing.T) {
	defer manager.clearAllListener()
	startListenerMonitor()
	syncChan1 := make(chan interface{})
	syncChan2 := make(chan interface{})
	manager.bindEventLister("msgType", func(newValue, oldValue interface{}) {
		time.Sleep(2 * time.Second)
		syncChan1 <- 1
	})

	manager.bindEventLister("msgType", func(newValue, oldValue interface{}) {
		if !strings.EqualFold(newValue.(string), "newValue") {
			t.Fatal("回调参数newvalue错误")
		}
		if !strings.EqualFold(oldValue.(string), "oldValue") {
			t.Fatal("回调参数oldValue错误")
		}
		syncChan2 <- 2
	})

	// send test msg
	manager.triggerEventListener(eventMsg{"msgType", "newValue", "oldValue"})

	isFirstArive := false
	isSecondArive := false

	select {
	case <-syncChan1:
		isFirstArive = true
		if !isSecondArive {
			t.Fatal("第一个监听器不应该先于第二个监听器执行")
		}
		break
	case <-syncChan2:
		isSecondArive = true
		if isFirstArive {
			t.Fatal("第一个监听器不应该先于第二个监听器执行")
		}
	case <-time.After(10 * time.Second):
		t.Fatal("监听消息处理超时")
		break
	}

}

// 测试同一个消息类型，有两个监听器的情况
func TestTriggerEventListener_onSameTypeMultiListener(t *testing.T) {
	defer manager.clearAllListener()
	startListenerMonitor()
	syncChan1 := make(chan interface{})
	manager.bindEventLister("msgType", func(newValue, oldValue interface{}) {
		syncChan1 <- 1
	})

	manager.bindEventLister("msgType", func(newValue, oldValue interface{}) {
		if !strings.EqualFold(newValue.(string), "newValue") {
			t.Fatal("回调参数newvalue错误")
		}
		if !strings.EqualFold(oldValue.(string), "oldValue") {
			t.Fatal("回调参数oldValue错误")
		}
		syncChan1 <- 2
	})

	// send test msg
	manager.triggerEventListener(eventMsg{"msgType", "newValue", "oldValue"})

	syncChan2 := make(chan interface{})
	go func() {
		count := 0
		for {
			<-syncChan1
			count = count + 1
			if count == 2 {
				break
			}
		}
		syncChan2 <- 1
	}()

	select {
	case <-syncChan2:
		break
	case <-time.After(10 * time.Second):
		t.Fatal("同一类型多个监听器，多个监听器没有同时执行")
		break
	}
}

// 测试同一个消息类型，有100个监听器的情况
func TestTriggerEventListener_onSameType100Listener(t *testing.T) {
	defer manager.clearAllListener()
	startListenerMonitor()
	syncChan1 := make(chan interface{})
	for i := 0; i < 100; i++ {
		manager.bindEventLister("msgType", func(newValue, oldValue interface{}) {
			syncChan1 <- i
		})
	}

	// send test msg
	manager.triggerEventListener(eventMsg{"msgType", "newValue", "oldValue"})

	syncChan2 := make(chan interface{})
	go func() {
		count := 0
		for {
			<-syncChan1
			count = count + 1
			if count == 100 {
				break
			}
		}
		syncChan2 <- 1
	}()

	select {
	case <-syncChan2:
		break
	case <-time.After(10 * time.Second):
		t.Fatal("同一类型100个监听器，100个监听器没有都触发执行")
		break
	}
}
