package main

import (
	"fmt"
	"github.com/bitxx/bitesla-strategy-exec/model"
)

//这是一个策略的测试

func Init() {
	fmt.Println("参数初始化")
}

func Run() {
	//test() //基本数据读取
	test1() //获取klines
	//test2() //获取trades
	//test3() //获取Depth
	//test4() //获取ticker
	//test5() //获取account
	//test6() //Orderplace
	//test7() //cancelOrder
	//test8() //获取oneOrder
	//test9() //获取UnfinishOrders
	//test10() //获取历史订单
	//test11() //获取交易所详情

}

func test() {
	//基本数据读取
	fmt.Println("currentUserID:", Srv.CurrentUserId)
	fmt.Println("exchangeId:", Srv.ExchangeId)
	fmt.Println("token:", Srv.Token)
	fmt.Println("appKey:", Srv.ApiKey)
	fmt.Println("appSecret:", Srv.ApiSecret)
}

func test1() {
	result, err := Srv.ExchangeGetKlineRecords(&model.ReqExchangeKline{ExchangeId: Srv.ExchangeId, Size: 10, Period: 0, Since: 10000, CurrencyPair: "BTC_USDT"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func test2() {
	result, err := Srv.ExchangeGetTrades(&model.ReqExchangeTrades{ExchangeId: Srv.ExchangeId, CurrencyPair: "BTC_USDT", Since: 1})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func test3() {
	result, err := Srv.ExchangeGetDepth(&model.ReqExchangeDepth{ExchangeId: Srv.ExchangeId, CurrencyPair: "BTC_USDT"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func test4() {
	result, err := Srv.ExchangeGetTicker(&model.ReqExchangeTicker{ExchangeId: Srv.ExchangeId, CurrencyPair: "BTC_USDT"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func test5() {
	result, err := Srv.ExchangeGetAccount(&model.ReqExchangeAccount{ExchangeId: Srv.ExchangeId, ApiKey: Srv.ApiKey, ApiSecret: Srv.ApiSecret})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func test6() {
	result, err := Srv.ExchangeOrderPlace(&model.ReqExchangeOrderPlace{ExchangeId: Srv.ExchangeId, ApiKey: Srv.ApiKey, ApiSecret: Srv.ApiSecret, Amount: "0", AccountType: 2, CurrencyPair: "BTC_USDT", OrderType: 2, Price: "0"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func test7() {
	result, err := Srv.ExchangeCancelOrder(&model.ReqExchangeCancelOrder{ExchangeId: Srv.ExchangeId, ApiKey: Srv.ApiKey, ApiSecret: Srv.ApiSecret, OrderId: "2xxx", CurrencyPair: "BTC_USDT"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func test8() {
	result, err := Srv.ExchangeGetOneOrder(&model.ReqExchangeOneOrder{ExchangeId: Srv.ExchangeId, ApiKey: Srv.ApiKey, ApiSecret: Srv.ApiSecret, OrderId: "2xxx", CurrencyPair: "BTC_USDT"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func test9() {
	result, err := Srv.ExchangeGetUnfinishOrders(&model.ReqExchangeUnfinishOrders{ExchangeId: Srv.ExchangeId, ApiKey: Srv.ApiKey, ApiSecret: Srv.ApiSecret, CurrencyPair: "BTC_USDT"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func test10() {
	result, err := Srv.ExchangeGetOrderHistorys(&model.ReqExchangeOrderHistory{ExchangeId: Srv.ExchangeId, ApiKey: Srv.ApiKey, ApiSecret: Srv.ApiSecret, CurrencyPair: "BTC_USDT", CurrentPage: 1, PageSize: 1})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func test11() {
	result, err := Srv.ExchangeGetExchangeDetail(&model.ReqExchangeDetail{ExchangeId: Srv.ExchangeId})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
