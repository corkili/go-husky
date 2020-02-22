package api

type RegisterReq struct {
	Username string `json:"username"`
	Phone string `json:"phone"`
	Password string `json:"password"`
}

type LoginReq struct {
	Phone string `json:"phone"`
	Password string `json:"password"`
}

type LogoutReq struct {
	Id uint `json:"id"`
}
