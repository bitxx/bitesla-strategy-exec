package model

type RespExchangeTrades struct {
	Code int     `json:"code"`
	Msg  string  `json:"msg"`
	Data []Trade `json:"data"`
}

type Trade struct {
	Tid          int64   `json:"tid"`
	Type         int32   `json:"type"`
	Amount       float64 `json:"amount"`
	Price        float64 `json:"price"`
	Date         int64   `json:"date"`
	CurrencyPair string  `json:"currencyPair"`
}
