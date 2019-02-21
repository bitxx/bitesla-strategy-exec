package service

import (
	"encoding/json"
	"github.com/jason-wj/bitesla-strategy-exec/model"
	"github.com/jason-wj/bitesla-strategy-exec/utils"
)

var (
	traderDetailUrl = "/trader/detail"
	traderStatusUrl = "/trader/status"
)

//GetTraderDetail
func (s *Service) TraderGetDetail(reqTraderDetail *model.ReqTraderDetail) (result *model.RespTraderDetail, err error) {
	h := utils.NewHttpSend(traderDetailUrl, utils.SendtypeJson)

	h.SetBody(reqTraderDetail)
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

//TraderUpdateStatus
func (s *Service) TraderUpdateStatus(reqTraderDetail *model.ReqTraderDetail) (result *model.RespTraderDetail, err error) {
	h := utils.NewHttpSend(traderStatusUrl, utils.SendtypeJson)

	h.SetBody(reqTraderDetail)
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
