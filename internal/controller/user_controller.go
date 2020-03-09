package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-husky/internal/api"
	"go-husky/internal/server"
	"go-husky/internal/service"
	"net/http"
	"sync"
)

type UserController struct {
	Controller
}

var userController *UserController
var userControllerOnce sync.Once

var userService = service.GetUserService()

var userRequestMappings = []server.RequestMapping{
	{
		UrlPath: "/user/register",
		Method: server.POST,
		Handler: UserRegister,
		Auth: false,
	},
	{
		UrlPath: "/user/login",
		Method: server.POST,
		Handler: UserLogin,
		Auth: false,
	},
	{
		UrlPath: "/user/logout",
		Method: server.POST,
		Handler: UserLogout,
		Auth: true,
	},
}

func GetUserController() (controller *UserController) {
	userControllerOnce.Do(func() {
		userController = &UserController{}
		userController.registerRequestMapping(userRequestMappings...)
	})
	return userController
}

func UserRegister(c *gin.Context) {
	var req = &api.RegisterReq{}
	var rsp *api.Response
	if err := c.ShouldBind(req); err != nil {
		logger.Error("register error: %s", err.Error())
		rsp = api.ReqDataInvalidResponse
	} else {
		rsp = userService.Register(req)
	}
	c.JSON(http.StatusOK, rsp)
}

func UserLogin(c *gin.Context) {
	var req = &api.LoginReq{}
	var rsp *api.Response
	if err := c.ShouldBind(req); err != nil {
		logger.Error("login error: %s", err.Error())
		rsp = api.ReqDataInvalidResponse
	} else {
		rsp = userService.Login(req)
		if rsp.Code == api.RspCodeSuccess {
			session := sessions.Default(c)
			id := rsp.Data["id"].(uint)
			session.Set("userId", id)
			err = session.Save()
			if err != nil {
				logger.Error("save session failed while login: %s", err.Error())
				rsp = api.SaveSessionErrorResponse
				userService.Logout(&api.LogoutReq{Id: id})
			}
		}
	}
	c.JSON(http.StatusOK, rsp)
}

func UserLogout(c *gin.Context) {
	var req = &api.LogoutReq{}
	var rsp *api.Response
	if err := c.ShouldBind(req); err != nil {
		rsp = api.ReqDataInvalidResponse
	} else {
		rsp = userService.Logout(req)
		if rsp.Code == api.RspCodeSuccess {
			session := sessions.Default(c)
			session.Delete("userId")
		}
	}
	c.JSON(http.StatusOK, rsp)
}

