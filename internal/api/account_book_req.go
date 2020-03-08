package api

type AccountBook struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}

type CreateOrUpdateAccountReq struct {
	Id uint `json:"id"`
	Datetime string `json:"datetime"`
	Type string `json:"type"`
	Money float64 `json:"money"`
	Description string `json:"description"`
	AccountBooks []*AccountBook `json:"accountBooks"`
}

type DeleteAccountReq struct {
	Ids []uint `json:"ids"`
}

type RetrieveAccountReq struct {
	All bool `json:"all"`
	Sort Sort `json:"sort"`
	Filters []Filter `json:"filters"`
	Pagination Pagination `json:"pagination"`
}

type RetrieveAccountBookReq struct {
	All bool `json:"all"`
}

type AccountStatisticReq struct {
	All bool `json:"all"`
}