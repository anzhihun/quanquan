/*
	消息派发器
*/
package event

// import (
// "code.huanshi.com/daca/common/log"
// "fmt"
// )

type eventMsg struct {
	msgType  string
	newValue interface{}
	oldValue interface{}
}

var eventChan chan eventMsg = make(chan eventMsg)
var isRunning bool = false

// 启动事件派发器
func RunEventDispather() {
	// 确保只启动一个派发线程
	if isRunning {
		return
	}
	isRunning = true

	go run()
	startListenerMonitor()
}

func run() {
	// 如果出现问题，继续执行，不要退出派发线程
	defer func() {
		if r := recover(); r != nil {
			// log.GetLogger().Warning(fmt.Sprintf("Event dispatcher recovered in %v", r))
		}
		run()
	}()

	for {
		event := <-eventChan
		manager.triggerEventListener(event)
	}
}
