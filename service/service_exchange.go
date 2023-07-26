package service

import (
	"encoding/json"
	"github.com/bitxx/bitesla-strategy-exec/model"
	"github.com/bitxx/bitesla-strategy-exec/utils"
)

var (
	exchangeOrderPlaceUrl        = rootUrl + "/exchange/orderPlace"
	exchangeCancelOrderUrl       = rootUrl + "/exchange/cancelOrder"
	exchangeGetOneOrderUrl       = rootUrl + "/exchange/getOneOrder"
	exchangeGetUnfinishOrdersUrl = rootUrl + "/exchange/getUnfinishOrders"
	exchangeGetOrderHistorysUrl  = rootUrl + "/exchange/getOrderHistorys"
	exchangeGetAccountUrl        = rootUrl + "/exchange/getAccount"
	exchangeGetTickerUrl         = rootUrl + "/exchange/getTicker"
	exchangeGetDepthUrl          = rootUrl + "/exchange/getDepth"
	exchangeGetKlineRecordsUrl   = rootUrl + "/exchange/getKlineRecords"
	exchangeGetTradesUrl         = rootUrl + "/exchange/getTrades"
	exchangeGetExchangeDetailUrl = rootUrl + "/exchange/getExchangeDetail"
	//exchangePutUrl               = rootUrl + "/exchange/put"
	//exchangeListUrl              = rootUrl + "/exchange/list"
)

// GetExchangeDetail
func (s *Service) ExchangeGetExchangeDetail(reqExchangeDetail *model.ReqExchangeDetail) (result interface{}, err error) {
	h := utils.NewHttpSend(exchangeGetExchangeDetailUrl, utils.SendtypeJson)

	h.SetBody(reqExchangeDetail)
	t := map[string]string{}
	t[token] = s.Token
	h.SetHeader(t)

	data, err := h.Post(false)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetOrderHistorys
func (s *Service) ExchangeGetOrderHistorys(reqOrderHistory *model.ReqExchangeOrderHistory) (result interface{}, err error) {
	h := utils.NewHttpSend(exchangeGetOrderHistorysUrl, utils.SendtypeJson)

	h.SetBody(reqOrderHistory)
	t := map[string]string{}
	t[token] = s.Token
	h.SetHeader(t)

	data, err := h.Post(false)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetUnfinishOrders
func (s *Service) ExchangeGetUnfinishOrders(reqUnfinishOrders *model.ReqExchangeUnfinishOrders) (result interface{}, err error) {
	h := utils.NewHttpSend(exchangeGetUnfinishOrdersUrl, utils.SendtypeJson)

	h.SetBody(reqUnfinishOrders)
	t := map[string]string{}
	t[token] = s.Token
	h.SetHeader(t)

	data, err := h.Post(false)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// CancelOrder
func (s *Service) ExchangeCancelOrder(reqCancelOrder *model.ReqExchangeCancelOrder) (result interface{}, err error) {
	h := utils.NewHttpSend(exchangeCancelOrderUrl, utils.SendtypeJson)

	h.SetBody(reqCancelOrder)
	t := map[string]string{}
	t[token] = s.Token
	h.SetHeader(t)

	data, err := h.Post(false)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetOneOrder
func (s *Service) ExchangeGetOneOrder(reqOneOrder *model.ReqExchangeOneOrder) (result interface{}, err error) {
	h := utils.NewHttpSend(exchangeGetOneOrderUrl, utils.SendtypeJson)

	h.SetBody(reqOneOrder)
	t := map[string]string{}
	t[token] = s.Token
	h.SetHeader(t)

	data, err := h.Post(false)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// OrderPlace
func (s *Service) ExchangeOrderPlace(reqOrderPlace *model.ReqExchangeOrderPlace) (result interface{}, err error) {
	h := utils.NewHttpSend(exchangeOrderPlaceUrl, utils.SendtypeJson)

	h.SetBody(reqOrderPlace)
	t := map[string]string{}
	t[token] = s.Token
	h.SetHeader(t)

	data, err := h.Post(false)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// 获取kline
func (s *Service) ExchangeGetKlineRecords(reqKline *model.ReqExchangeKline) (result *model.RespExchangeKlines, err error) {
	h := utils.NewHttpSend(exchangeGetKlineRecordsUrl, utils.SendtypeJson)

	h.SetBody(reqKline)
	t := map[string]string{}
	t[token] = s.Token
	h.SetHeader(t)

	data, err := h.Post(false)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// 获取Trades
func (s *Service) ExchangeGetTrades(reqTrades *model.ReqExchangeTrades) (result *model.RespExchangeKlines, err error) {
	h := utils.NewHttpSend(exchangeGetTradesUrl, utils.SendtypeJson)

	h.SetBody(reqTrades)
	t := map[string]string{}
	t[token] = s.Token
	h.SetHeader(t)

	data, err := h.Post(false)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// 获取Depths
func (s *Service) ExchangeGetDepth(reqDepth *model.ReqExchangeDepth) (result *model.RespExchangeDepth, err error) {
	h := utils.NewHttpSend(exchangeGetDepthUrl, utils.SendtypeJson)

	h.SetBody(reqDepth)
	t := map[string]string{}
	t[token] = s.Token
	h.SetHeader(t)

	data, err := h.Post(false)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// 获取Ticker
func (s *Service) ExchangeGetTicker(reqTicker *model.ReqExchangeTicker) (result *model.RespExchangeTicker, err error) {
	h := utils.NewHttpSend(exchangeGetTickerUrl, utils.SendtypeJson)

	h.SetBody(reqTicker)
	t := map[string]string{}
	t[token] = s.Token
	h.SetHeader(t)

	data, err := h.Post(false)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// 获取Account
func (s *Service) ExchangeGetAccount(reqAccount *model.ReqExchangeAccount) (result *model.RespExchangeAccounts, err error) {
	h := utils.NewHttpSend(exchangeGetAccountUrl, utils.SendtypeJson)

	h.SetBody(reqAccount)
	t := map[string]string{}
	t[token] = s.Token
	h.SetHeader(t)

	data, err := h.Post(false)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}
