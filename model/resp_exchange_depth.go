package model

type RespExchangeDepth struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data Depth  `json:"data"`
}

type Depth struct {
	AskList DepthRecords `json:"AskList"`
	BidList DepthRecords `json:"BidList"`
}

type DepthRecords struct {
	List []DepthRecord `json:"list"`
}

type DepthRecord struct {
	Price  float64 `json:"price"`
	Amount float64 `json:"amount"`
}
