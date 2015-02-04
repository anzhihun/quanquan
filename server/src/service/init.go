package service

func Init() {
	initUsers()
	go commServer.start()
	commServer.broadcastMe()
	listenViewMsg()
}
