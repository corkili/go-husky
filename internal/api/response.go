package api

type Response struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	UiMsg string `json:"uiMsg"`
	Data map[string]interface{} `json:"data"`
}

var NoAuthResponse = &Response{
	Code: RspCodeNoAuth,
	Msg:  "no auth, login first",
	UiMsg: "请先登录",
	Data: nil,
}

var SuccessWithoutData = &Response{
	Code: RspCodeSuccess,
	Msg:  "success",
	UiMsg: "操作成功",
	Data: nil,
}

var ReqDataInvalidResponse = &Response{
	Code: RspCodeReqDataInvalid,
	Msg:  "request data is invalid",
	UiMsg: "请求异常，请重试",
	Data: nil,
}

var SaveSessionErrorResponse = &Response{
	Code: RspCodeSaveSessionError,
	Msg:  "login with error: save session failed",
	UiMsg: "服务器异常，请联系管理员",
	Data: nil,
}