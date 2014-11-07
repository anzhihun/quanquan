// 定义事件的相关接口
package event

import (
	"sync"
)

var eventChanMutex sync.Mutex

// 触发消息
// msgType自定义，发送方和接收方一致就可以了。
func Trigger(msgType string, newValue, oldValue interface{}) {
	eventChanMutex.Lock()
	defer eventChanMutex.Unlock()

	eventChan <- eventMsg{msgType, newValue, oldValue}
}

// 注册消息监听，返回一个id号，通过该id号可以通过off接口注销该监听
func On(msgType string, eventFunc func(interface{}, interface{})) int32 {
	return manager.bindEventLister(msgType, eventFunc)
}

// 取消消息监听
func Off(id int32) {
	manager.unbindEventLister(id)
}
