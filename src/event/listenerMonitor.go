//所有时间监听器的监测器，监控所有的监听器执行情况
package event

import (
	// "code.huanshi.com/daca/common/log"
	// "fmt"
	"strconv"
	"time"
)

type monitorListener struct {
	info      string
	startTime int64
	syncChan  chan int64
}

var monitorChan chan monitorListener = make(chan monitorListener)

func startListenerMonitor() {
	go func() {

		// 如果出现问题，继续执行，不要退出派发线程
		defer func() {
			if r := recover(); r != nil {
				// log.GetLogger().Warning(fmt.Sprintf("Event Listener Monitor recovered in %v", r))
			}
			runMonitor()
		}()

		runMonitor()

	}()
}

func runMonitor() {
	for {
		listener := <-monitorChan
		go func() {
			select {
			case <-listener.syncChan:
				// log.GetLogger().Debug(fmt.Sprintf("%s execute use %d ns", listener.info, endTime-listener.startTime))
			case <-time.After(1 * time.Minute):
				// log.GetLogger().Warning(listener.info + " execute more than 1 minute!")
			}
		}()
	}
}

func addListener2Monitor(l *eventListener, syncChan chan int64) {
	listener := monitorListener{"msgType " + l.eventType + " 's listener(id=" + strconv.FormatInt(int64(l.id), 10) + ")", time.Now().UnixNano(), syncChan}
	monitorChan <- listener
}
