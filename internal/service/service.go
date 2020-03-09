package service

import (
	"go-husky/internal/api"
	"go-husky/internal/log"
)

var logger = log.GetLogger()

func BuildResponse(code int, msg string, uiMsg string, data map[string]interface{}) *api.Response {
	return &api.Response{
		Code:  code,
		Msg:   msg,
		UiMsg: uiMsg,
		Data:  data,
	}
}
