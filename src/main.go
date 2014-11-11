// start a http server to comunicate with UI

package main

import (
	"code.google.com/p/go.net/websocket"
	"controller"
	"encoding/json"
	"event"
	"fmt"
	"github.com/gocraft/web"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"service"
	"user"
)

type Context struct {
}

func (this *Context) getIndex(rw web.ResponseWriter, req *web.Request) {
	// return index.html as index page
	rootDir, _ := os.Getwd()
	pathSeparator := string(filepath.Separator)
	indexFileName := rootDir + pathSeparator + "www" + pathSeparator + "index.html"
	indexContent, _ := ioutil.ReadFile(indexFileName)
	rw.Write(indexContent)
}

func (this *Context) getUsers(rw web.ResponseWriter, req *web.Request) {
	users := user.UserManager.AllUser()
	contents, err := json.Marshal(users)
	if err != nil {
		rw.Write([]byte("error: " + err.Error()))
	} else {
		rw.Write(contents)
	}
	// req.ParseForm()
}

func (this *Context) onWsConnection(ws *websocket.Conn) {
	// return index.html as index page

	fmt.Println("receive ws connect")
}

func main() {

	event.RunEventDispather()
	service.Init()
	controller.Init()
	startHttpServer()
}

func startHttpServer() {
	router := web.New(Context{}). // Create your router
					Middleware(web.LoggerMiddleware). // Use some included middleware
					Middleware(web.ShowErrorsMiddleware)

	rootDir, _ := os.Getwd()
	rootDir = rootDir + string(filepath.Separator) + "www"
	router.Middleware(web.StaticMiddleware(rootDir))
	router.Middleware(controller.WebSocketMiddleware)

	router.Get("/", (*Context).getIndex)
	router.Get("/users", (*Context).getUsers)

	http.ListenAndServe("localhost:53240", router) // Start the server!
}
