// 管理各种消息
package event

import (
	// "code.huanshi.com/daca/common/log"
	// "code.huanshi.com/daca/common/ultils"
	"sync"
	// "time"
	// "fmt"
)

type eventManager struct {
	idGenerator int32
	// key 为消息类型，value为监听器的map，key为监听器id，value为消息回调方法
	listeners map[string][]*eventListener
	rwLocker  sync.RWMutex
}

var manager *eventManager = &eventManager{idGenerator: 1, listeners: make(map[string][]*eventListener)}

func (self *eventManager) getListenerCount(msgType string) int {
	self.rwLocker.RLock()
	self.rwLocker.RUnlock()

	return len(self.listeners[msgType])
}

func (self *eventManager) clearAllListener() {
	self.rwLocker.Lock()
	self.rwLocker.Unlock()

	for msgType := range self.listeners {
		if self.listeners[msgType] != nil {
			for _, listener := range self.listeners[msgType] {
				listener.StopProcess()
			}
		}
		self.listeners[msgType] = nil
		delete(self.listeners, msgType)
	}
}

func (self *eventManager) unbindEventLister(id int32) {
	self.rwLocker.Lock()
	defer self.rwLocker.Unlock()

	for msgType := range self.listeners {
		listeners := self.listeners[msgType]
		for index, listener := range listeners {
			if listener.id == id {
				listener.StopProcess()
				listeners := append(listeners[:index], listeners[index+1:]...)
				// ultils.SliceHelper(&listeners).Remove(index)
				self.listeners[msgType] = listeners
				break
			}
		}

		if len(self.listeners[msgType]) == 0 {
			delete(self.listeners, msgType)
		}
	}
}

func (self *eventManager) bindEventLister(msgType string, eventFunc func(interface{}, interface{})) int32 {
	self.rwLocker.Lock()
	defer self.rwLocker.Unlock()

	self.idGenerator = self.idGenerator + 1
	var locker sync.Mutex
	listener := &eventListener{eventType: msgType, id: self.idGenerator, eventQueue: make([]eventMsg, 0), eventFunc: eventFunc, cond: sync.NewCond(&locker)}
	self.listeners[msgType] = append(self.listeners[msgType], listener)
	self.listeners[msgType][len(self.listeners[msgType])-1].listen()

	// log.GetLogger().Debug("bind listener on " + msgType)

	return self.idGenerator
}

func (self *eventManager) triggerEventListener(event eventMsg) {
	self.rwLocker.RLock()
	defer self.rwLocker.RUnlock()

	if self.listeners[event.msgType] == nil {
		return
	}

	for index, _ := range self.listeners[event.msgType] {
		self.listeners[event.msgType][index].onEvent(event)
	}
}
