package api

type Pagination struct {
	CurrentPage uint `json:"currentPage"`
	PageSize uint `json:"pageSize"`
	TotalCount uint `json:"totalCount"`
}

type SortRule struct {
	Field string `json:"field"`
	Method string `json:"method"`
}

type Sort struct {
	SortRules []SortRule `json:"sortRules"`
}

type Filter struct {
	Field string `json:"field"`
	Operation string `json:"operation"`
	Value interface{} `json:"value"`
}