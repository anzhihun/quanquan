package service

func Init() {
	initUsers()
	initChannels()

	go commServer.start()
	commServer.broadcastMe()
	listenViewMsg()
}
