package model

type RespTraderDetail struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data Trader `json:"data"`
}

type Trader struct {
	TraderId    int64  `json:"traderId" example:"948904443912"`
	StrategyId  int64  `json:"strategyId" example:"23575003451411"`
	ExchangeId  int64  `json:"exchangeId" example:"2358885120275906"`
	Name        string `json:"name"`
	Status      int32  `json:"status"`
	Description string `json:"description"`
	CreateTime  int64  `json:"createTime"`
	UpdateTime  int64  `json:"updateTime"`
}
