package controller

import (
	"github.com/gocraft/web"
	"os"
	"path/filepath"
)

type RootContext struct {
}

type UserContext struct {
	*RootContext
}

type ChannelContext struct {
	*RootContext
}

func Init() *web.Router {

	rtMsgController.listenToViewMessage()

	rootRouter := web.New(RootContext{}). // Create your rootRouter
						Middleware(web.LoggerMiddleware). // Use some included middleware
						Middleware(web.ShowErrorsMiddleware)

	// handle index request
	rootRouter.Get("/", (*RootContext).GetIndex)

	// handle static file requests
	rootDir, _ := os.Getwd()
	rootDir = rootDir + string(filepath.Separator) + "www"
	rootRouter.Middleware(web.StaticMiddleware(rootDir))

	// handle websocket requests
	rootRouter.Middleware(webSocketMiddleware)

	// handle user requests
	userRouter := rootRouter.Subrouter(UserContext{}, "/user")
	userRouter.Get("/", (*UserContext).GetUsers)
	userRouter.Post("/login", (*UserContext).Login)
	userRouter.Post("/signup", (*UserContext).SignUp)

	// handle channel requests
	channelRouter := rootRouter.Subrouter(ChannelContext{}, "/channel")
	channelRouter.Post("/", (*ChannelContext).addChannel)

	return rootRouter
}
