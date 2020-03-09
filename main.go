package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"go-husky/internal/controller"
	"go-husky/internal/log"
	"go-husky/internal/server"
	"go-husky/internal/service"
)

var logger = log.GetLogger()

func main() {

	store := cookie.NewStore([]byte("secret"))

	builder := server.Builder{}
	builder.SetPort(8000).
		EnableLog(true).
		EnableRecovery(true).
		SetAuthHandler(service.CheckIfUserLogin).
		AddMiddleware(sessions.Sessions("auth", store)).
		AddRequestMapping(controller.GetUserController().GetRequestMappings()...).
		AddRequestMapping(controller.GetAccountBookController().GetRequestMappings()...)
	ginServer := builder.Build()
	err := ginServer.Start()
	if err != nil {
		logger.Error(err)
	}

}
