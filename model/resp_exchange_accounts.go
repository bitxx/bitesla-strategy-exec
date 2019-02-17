package model

type RespExchangeAccounts struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data []Account `json:"data"`
}

type Account struct {
	Asset       float64                `json:"asset"`
	EtAsset     float64                `json:"etAsset"`
	SubAccounts map[string]*SubAccount `json:"subAccounts"`
}

type SubAccount struct {
	Currency     string  `json:"currency"`
	Amount       float64 `json:"amount"`
	ForzenAmount float64 `json:"forzenAmount"`
	LoanAmount   float64 `json:"loanAmount"`
}
