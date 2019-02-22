package model

type RespTraderStatus struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data Trader `json:"data"`
}
