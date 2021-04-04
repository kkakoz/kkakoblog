package model

type Pager struct {
	Page     int `form:"page"`
	PageSize int `form:"page_size"`
}

type Res struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Count int64       `json:"count"`
}
