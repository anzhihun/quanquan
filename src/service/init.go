package service

func Init() {
	go commServer.start()
	commServer.broadcastMe()
	listenViewMsg()
}
