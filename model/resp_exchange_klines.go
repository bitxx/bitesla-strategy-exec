package model

type RespExchangeKlines struct {
	Code int     `json:"code"`
	Msg  string  `json:"msg"`
	Data []Kline `json:"data"`
}

type Kline struct {
	CurrencyPair string  `json:"currencyPair"`
	Timestamp    int64   `json:"timestamp"`
	Open         float64 `json:"open"`
	Close        float64 `json:"close"`
	High         float64 `json:"high"`
	Low          float64 `json:"low"`
	Vol          float64 `json:"vol"`
}
