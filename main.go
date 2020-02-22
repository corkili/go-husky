package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
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
		AddRequestMapping(controller.GetUserController().GetRequestMappings()...)
	ginServer := builder.Build()
	err := ginServer.Start()
	if err != nil {
		logger.Error(err)
	}

}

func testSession(c *gin.Context) {
	session := sessions.Default(c)
	var count int
	v := session.Get("count")
	if v == nil {
		count = 0
	} else {
		count = v.(int)
		count++
	}
	session.Set("count", count)
	session.Save()
	c.JSON(200, gin.H{"count": count})
}
