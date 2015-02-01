package controller

import (
	"github.com/gocraft/web"
	"utils"
)

type RootContext struct {
}

func Init() *web.Router {

	rtMsgController.listenToViewMessage()

	rootRouter := web.New(RootContext{}).
		Middleware(web.LoggerMiddleware).
		Middleware(web.ShowErrorsMiddleware)

	// handle index request
	// rootRouter.Get("/", (*RootContext).GetIndex)

	// handle static file requests
	rootDir := utils.MakePath([]string{utils.GetRootDir(), "www"})
	rootRouter.Middleware(web.StaticMiddleware(rootDir))

	// handle websocket requests
	rootRouter.Middleware(webSocketMiddleware)

	// route user requests
	initUserRouter(rootRouter)

	// route channel requests
	initChannelRouter(rootRouter)

	// route resource requests
	initResRouter(rootRouter)

	return rootRouter
}

func initUserRouter(rootRouter *web.Router) {
	userRouter := rootRouter.Subrouter(UserContext{}, "/user")
	userRouter.Get("/", (*UserContext).GetUsers)
	userRouter.Post("/login", (*UserContext).Login)
	userRouter.Post("/signup", (*UserContext).SignUp)
}

func initChannelRouter(rootRouter *web.Router) {
	channelRouter := rootRouter.Subrouter(ChannelContext{}, "/channel")
	channelRouter.Post("/", (*ChannelContext).addChannel)
	channelRouter.Get("/", (*ChannelContext).getChannels)
	channelRouter.Post("/inviteUser", (*ChannelContext).inviteUserToChannel)
}

func initResRouter(rootRouter *web.Router) {
	resRouter := rootRouter.Subrouter(resController{}, "/res")
	resRouter.Get("/language", (*resController).getLanguage)
}
