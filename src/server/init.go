package server

func Init() {
	go commServer.start()
	commServer.broadcastMe()
}
