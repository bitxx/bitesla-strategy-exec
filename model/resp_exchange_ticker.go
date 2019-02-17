package model

type RespExchangeTicker struct {
	Code int     `json:"code"`
	Msg  string  `json:"msg"`
	Data Ticker `json:"data"`
}

type Ticker struct {
	ContractType string  `json:"contractType"`
	Pair         string  `json:"pair"`
	Last         float64 `json:"last"`
	Buy          float64 `json:"buy"`
	Sell         float64 `json:"sell"`
	High         float64 `json:"high"`
	Low          float64 `json:"low"`
	Vol          float64 `json:"vol"`
	Date         uint64  `json:"date"`
}