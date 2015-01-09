// start a http server to comunicate with UI

package main

import (
	"controller"
	"event"
	"net/http"
	"service"
	"user"
)

func main() {
	event.RunEventDispather()
	user.Init()
	service.Init()
	startHttpServer()
}

func startHttpServer() {

	http.ListenAndServe("localhost:52013", controller.Init()) // Start the server!
}
