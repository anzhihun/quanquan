// start a http server to comunicate with UI

package main

import (
	"controller"
	"encoding/json"
	"event"
	// "fmt"
	"github.com/gocraft/web"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"service"
	"user"
	"utils"
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

func (this *Context) addChannel(rw web.ResponseWriter, req *web.Request) {
	result, _ := ioutil.ReadAll(req.Body)
	req.Body.Close()
	params, _ := utils.DecodeJsonMsg(string(result))
	newChannel := user.Channel{params["name"].(string), []*user.User{}}
	user.AddChannel(&newChannel)
}

func (this *Context) login(rw web.ResponseWriter, req *web.Request) {
	result, _ := ioutil.ReadAll(req.Body)
	req.Body.Close()
	params, err := utils.DecodeJsonMsg(string(result))
	if err != nil {
		http.Error(rw, err.Error(), 404)
		return
	}

	if !user.Validate(params["name"].(string), params["password"].(string)) {
		http.Error(rw, "invalid user name or password", 500)
	}
}

func (this *Context) signUp(rw web.ResponseWriter, req *web.Request) {
	result, _ := ioutil.ReadAll(req.Body)
	req.Body.Close()
	params, err := utils.DecodeJsonMsg(string(result))
	if err != nil {
		http.Error(rw, err.Error(), 404)
		return
	}

	if err = user.SignUp(params["name"].(string), params["password"].(string)); err != nil {
		http.Error(rw, err.Error(), 500)
	}
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
	router.Post("/channel", (*Context).addChannel)
	router.Post("/login", (*Context).login)
	router.Post("/signup", (*Context).signUp)

	http.ListenAndServe("localhost:53240", router) // Start the server!
}
