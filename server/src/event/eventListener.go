package event

import (
	// "code.huanshi.com/daca/common/log"
	// "fmt"
	// "strconv"
	"sync"
	"time"
)

type eventListener struct {
	eventType  string
	id         int32
	eventQueue []eventMsg
	cond       *sync.Cond
	eventFunc  func(interface{}, interface{})
	stop       bool
}

// 监听事件
func (self *eventListener) listen() {
	go self.process()
}

// 处理接收到的事件
func (self *eventListener) process() {
	for {
		self.cond.L.Lock()

		// wait for event
		for len(self.eventQueue) == 0 {
			// stop goroutine
			if self.stop {
				// log.GetLogger().Info(fmt.Sprintf("%s listener %d stop! left event len = %d", self.eventType, self.id, len(self.eventQueue)))
				self.stop = false
				self.cond.L.Unlock()
				return
			}
			self.cond.Wait()
		}

		// make a copy to lock less time
		copyEvents := []eventMsg{}
		copyEvents = append(copyEvents, self.eventQueue...)
		self.eventQueue = self.eventQueue[:0]
		self.cond.L.Unlock()

		for _, eventMsg := range copyEvents {
			self.executeEventFunc(eventMsg)
		}
	}
}

// 执行事件响应函数，同时避免一次事件执行，导致整个事件处理循环终止
func (self *eventListener) executeEventFunc(event eventMsg) {
	defer func() {
		if r := recover(); r != nil {
			// log.GetLogger().Error(fmt.Sprintf(" %s event listener %d process recovered in %v", self.eventType, self.id, r))
		}
	}()

	// 添加时间响应函数执行时间监测
	//log.GetLogger().Debug(fmt.Sprintf(" exec msg type: %s, value: %v", event.msgType, event.newValue))
	syncChan := make(chan int64)
	addListener2Monitor(self, syncChan)
	self.eventFunc(event.newValue, event.oldValue)
	syncChan <- time.Now().UnixNano()
}

func (self *eventListener) onEvent(event eventMsg) {
	if self.isEventsFull() {
		// log.GetLogger().Warning(self.eventType + " listener's queue is full! id = " + strconv.FormatInt(int64(self.id), 10) + ", discard new event.")
		return
	}

	self.cond.L.Lock()
	defer self.cond.L.Unlock()

	//log.GetLogger().Debug(fmt.Sprintf(" receive msg type: %s, value: %v", event.msgType, event.newValue))
	self.eventQueue = append(self.eventQueue, event)
	self.cond.Signal()
}

func (self *eventListener) isEventsFull() bool {
	self.cond.L.Lock()
	defer self.cond.L.Unlock()
	return len(self.eventQueue) >= 1000
}

func (self *eventListener) StopProcess() {
	self.cond.L.Lock()
	defer self.cond.L.Unlock()
	self.stop = true
	self.cond.Signal()
}
