package controller

import (
	"github.com/gin-gonic/gin"
	"go-husky/internal/api"
	"go-husky/internal/server"
	"go-husky/internal/service"
	"net/http"
	"sync"
)

type AccountBookController struct {
	Controller
}

var accountBookController *AccountBookController
var accountBookControllerOnce sync.Once

var accountBookService = service.GetAccountBookService()

var accountBookRequestMappings = []server.RequestMapping{
	{
		UrlPath: "/account_book/create_account",
		Method: server.POST,
		Handler: CreateAccount,
		Auth: true,
	},
	{
		UrlPath: "/account_book/update_account",
		Method: server.POST,
		Handler: UpdateAccount,
		Auth: true,
	},
	{
		UrlPath: "/account_book/retrieve_account",
		Method: server.POST,
		Handler: RetrieveAccount,
		Auth: true,
	},
	{
		UrlPath: "/account_book/delete_account",
		Method: server.POST,
		Handler: DeleteAccount,
		Auth: true,
	},
	{
		UrlPath: "/account_book/retrieve_account_book",
		Method: server.POST,
		Handler: RetrieveAccountBook,
		Auth: true,
	},
}

func GetAccountBookController() (controller *AccountBookController) {
	accountBookControllerOnce.Do(func() {
		accountBookController = &AccountBookController{}
		accountBookController.registerRequestMapping(accountBookRequestMappings...)
	})
	return accountBookController
}

func CreateAccount(c *gin.Context) {
	var req = &api.CreateOrUpdateAccountReq{}
	var rsp *api.Response
	if err := c.ShouldBind(req); err != nil {
		logger.Error("create account error: %s", err.Error())
		rsp = api.ReqDataInvalidResponse
	} else {
		req.Id = 0
		user, err := userService.GetCurrentUser(c)
		if err != nil {
			logger.Error(err.Error())
			rsp = api.NoAuthResponse
		} else {
			rsp = accountBookService.SaveAccount(req, user)
		}
	}
	c.JSON(http.StatusOK, rsp)
}

func UpdateAccount(c *gin.Context) {
	var req = &api.CreateOrUpdateAccountReq{}
	var rsp *api.Response
	if err := c.ShouldBind(req); err != nil {
		logger.Error("create account error: %s", err.Error())
		rsp = api.ReqDataInvalidResponse
	} else {
		if req.Id == 0 {
			rsp = service.BuildResponse(api.RspCodeAccountBookErrorAccountIdEmpty,
				"account id is zero", "更新失败，请重试", nil)
		} else {
			user, err := userService.GetCurrentUser(c)
			if err != nil {
				logger.Error(err.Error())
				rsp = api.NoAuthResponse
			} else {
				rsp = accountBookService.SaveAccount(req, user)
			}
		}
	}
	c.JSON(http.StatusOK, rsp)
}

func RetrieveAccount(c *gin.Context) {
	var req = &api.RetrieveAccountReq{}
	var rsp *api.Response
	if err := c.ShouldBind(req); err != nil {
		logger.Error("register error: %s", err.Error())
		rsp = api.ReqDataInvalidResponse
	} else {
		user, err := userService.GetCurrentUser(c)
		if err != nil {
			logger.Error(err.Error())
			rsp = api.NoAuthResponse
		} else {
			rsp = accountBookService.RetrieveAccount(req, user)
		}
	}
	c.JSON(http.StatusOK, rsp)
}

func RetrieveAccountBook(c *gin.Context) {
	var req = &api.RetrieveAccountBookReq{}
	var rsp *api.Response
	if err := c.ShouldBind(req); err != nil {
		logger.Error("register error: %s", err.Error())
		rsp = api.ReqDataInvalidResponse
	} else {
		user, err := userService.GetCurrentUser(c)
		if err != nil {
			logger.Error(err.Error())
			rsp = api.NoAuthResponse
		} else {
			rsp = accountBookService.RetrieveAccountBook(req, user)
		}
	}
	c.JSON(http.StatusOK, rsp)
}

func DeleteAccount(c *gin.Context) {
	var req = &api.DeleteAccountReq{}
	var rsp *api.Response
	if err := c.ShouldBind(req); err != nil {
		logger.Error("register error: %s", err.Error())
		rsp = api.ReqDataInvalidResponse
	} else {
		user, err := userService.GetCurrentUser(c)
		if err != nil {
			logger.Error(err.Error())
			rsp = api.NoAuthResponse
		} else {
			rsp = accountBookService.DeleteAccount(req, user)
		}
	}
	c.JSON(http.StatusOK, rsp)
}

