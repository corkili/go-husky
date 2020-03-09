package service

import (
	"errors"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-husky/internal/api"
	"go-husky/internal/common"
	"go-husky/internal/dao"
	"go-husky/internal/entity"
	"sync"
)

type UserService struct {
	userDao *dao.UserDao
	onlineUsers map[uint]*entity.User
}

var userService *UserService
var userServiceOnce sync.Once

func GetUserService() (controller *UserService) {
	userServiceOnce.Do(func() {
		userService = &UserService{}
		userService.onlineUsers = make(map[uint]*entity.User)
		userService.userDao = dao.GetUserDao()
	})
	return userService
}

func CheckIfUserLogin(c *gin.Context) (ok bool, rsp *api.Response) {
	session := sessions.Default(c)
	userIdRes := session.Get("userId")
	var userId uint = 0
	if userIdRes != nil {
		userId = userIdRes.(uint)
	}
	service := GetUserService()
	_, exist := service.onlineUsers[userId]
	if exist {
		return true, nil
	} else {
		return false, api.NoAuthResponse
	}
}

func (s *UserService) CheckIfUserLogin(user *entity.User) (ok bool, rsp *api.Response) {
	if user != nil {
		_, exist := s.onlineUsers[user.ID]
		if exist {
			return true, nil
		} else {
			return false, api.NoAuthResponse
		}
	} else {
		return false, api.NoAuthResponse
	}
}

func (s *UserService) GetCurrentUser(c *gin.Context) (user *entity.User, err error) {
	session := sessions.Default(c)
	userIdRes := session.Get("userId")
	var userId uint = 0
	if userIdRes != nil {
		userId = userIdRes.(uint)
	}
	user, exist := s.onlineUsers[userId]
	if exist {
		return user, nil
	} else {
		return nil, errors.New(api.NoAuthResponse.Msg)
	}
}

func (s *UserService) Register(req *api.RegisterReq) *api.Response {
	var code int
	var msg string
	var uiMsg string
	var data map[string]interface{}
	var username = req.Username
	var phone = req.Phone
	var password = req.Password
	if username == "" {
		code, msg, uiMsg = api.RspCodeUserErrorUsernameEmpty, "username is empty", "用户名不能为空"
	} else if phone == "" {
		code, msg, uiMsg = api.RspCodeUserErrorPhoneEmpty, "phone is empty", "手机号不能为空"
	} else if password == "" {
		code, msg, uiMsg = api.RspCodeUserErrorPasswordEmpty, "password is empty", "密码不能为空"
	} else {
		_, exist, err := s.userDao.FindByPhone(phone)
		if err != nil {
			code, msg, uiMsg = api.RspCodeDbError, "query db failed", "查询数据异常"
		} else if exist {
			code, msg, uiMsg = api.RspCodeUserErrorPhoneExists, fmt.Sprintf("phone[%s] already exists", phone), "手机号已存在"
		} else {
			user := entity.User {
				Phone:    phone,
				Password: common.GetSecretPassword(password),
				Username: username,
				Roll:     entity.USER,
			}
			err = s.userDao.CreateEntity(&user)
			if err != nil {
				logger.Error("register user error occurs in dao: %s", err.Error())
				code, msg, uiMsg = api.RspCodeDbError, "insert db failed", "数据操作异常"
			} else {
				code, msg, uiMsg = api.RspCodeSuccess, "success", "注册成功"
				data = map[string]interface{}{
					"id": user.ID,
					"phone": phone,
					"username": username,
				}
			}
		}
	}
	return &api.Response{
		Code: code,
		Msg:  msg,
		UiMsg: uiMsg,
		Data: data,
	}
}

func (s *UserService) Login(req *api.LoginReq) *api.Response {
	var code int
	var msg string
	var uiMsg string
	var data map[string]interface{}
	var phone = req.Phone
	var password = req.Password
	if phone == "" {
		code, msg, uiMsg = api.RspCodeUserErrorPhoneEmpty, "phone is empty", "手机号不能为空"
	} else if password == "" {
		code, msg, uiMsg = api.RspCodeUserErrorPasswordEmpty, "password is empty", "密码不能为空"
	} else {
		user, exist, err := s.userDao.FindByPhone(phone)
		if err != nil {
			code, msg, uiMsg = api.RspCodeDbError, "query db failed", "数据查询异常"
		} else if !exist {
			code, msg, uiMsg = api.RspCodeUserErrorPhoneNoExists, fmt.Sprintf("phone[%s] no exists", phone), "手机号不存在"
		} else if !common.ValidatePassword(password, user.Password) {
			code, msg, uiMsg = api.RspCodeUserErrorPasswordInvalid, fmt.Sprintf("invalid password"), "密码错误"
		} else {
			s.onlineUsers[user.ID] = user
			code, msg, uiMsg = api.RspCodeSuccess, "success", "登录成功"
			data = map[string]interface{}{
				"id": user.ID,
				"phone": user.Phone,
				"username": user.Username,
			}
		}
	}

	return &api.Response{
		Code: code,
		Msg:  msg,
		UiMsg: uiMsg,
		Data: data,
	}
}

func (s *UserService) Logout(req *api.LogoutReq) *api.Response {
	_, exist := s.onlineUsers[req.Id]
	if exist {
		delete(s.onlineUsers, req.Id)
	}
	return api.SuccessWithoutData
}




