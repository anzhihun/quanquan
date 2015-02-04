// start a http server to comunicate with UI

package main

import (
	"controller"
	"event"
	"net/http"
	"service"
)

func main() {
	event.RunEventDispather()
	service.Init()
	startHttpServer()
}

func startHttpServer() {

	http.ListenAndServe(":52013", controller.Init()) // Start the server!
}
