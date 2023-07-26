package main

import (
	"errors"
	"fmt"
	"github.com/bitxx/bitesla-strategy-exec/model"
	"github.com/bitxx/bitesla-strategy-exec/service"
	"github.com/henrylee2cn/goutil/calendar/cron"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

//测试命令： go run *.go 1 2358885120275906 948904443912 23575003451411 eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImlkZWFfd2pAMTYzLmNvbSIsInBhc3N3b3JkIjoic2phMTIzIiwidXNlcklkIjoxLCJleHAiOjE1NTEzMjQzNzUsImlzcyI6Ind1aiJ9.U1dI1woAgclu0tP18JI0xBMzVHCq5tGugPi7OiSftuY cf998e2f-ed32b3f8-42c02442-1011e 91e6e532-3bb37cad-c0a689ca-75ee5

var Srv *service.Service

// 指定时间更新一次,10分钟
const Timer string = "@every 10s"

const (
	statusStop  int32 = 0
	statusStart int32 = 1
)

var (
	c         *cron.Cron
	s         chan os.Signal
	startTime int64
)

func init() {
	Init()
}

func main() {
	if len(os.Args) != 8 {
		panic(errors.New("参数个数错误"))
	}

	//当前用户id 1
	currentUserIdStr := os.Args[1]
	//当前交易所 2358885120275906
	exchangeIdStr := os.Args[2]
	//交易执行 948904443912
	traderIdStr := os.Args[3]
	//23575003451411
	strategyIdStr := os.Args[4]
	//令牌
	token := os.Args[5]
	//密钥key cf998e2f-ed32b3f8-42c02442-1011e
	apiKey := os.Args[6]
	//密钥secret 91e6e532-3bb37cad-c0a689ca-75ee5
	apiSecret := os.Args[7]

	currentUserId, _ := strconv.ParseInt(currentUserIdStr, 10, 64)
	exchangeId, _ := strconv.ParseInt(exchangeIdStr, 10, 64)
	traderId, _ := strconv.ParseInt(traderIdStr, 10, 64)
	strategyId, _ := strconv.ParseInt(strategyIdStr, 10, 64)
	//初始化服务
	Srv = &service.Service{ApiKey: apiKey, ApiSecret: apiSecret, CurrentUserId: currentUserId, Token: token, ExchangeId: exchangeId, TraderId: traderId, StrategyId: strategyId}

	c = cron.New()
	startTime = time.Now().Unix() //策略启动时间
	err := c.AddFunc(Timer, task) //定时估算任务
	if err != nil {
		fmt.Println("策略启动异常：" + err.Error())
		return
	}
	c.Start()

	//持续运行
	s = make(chan os.Signal)
	signal.Notify(s, syscall.SIGINT)
	<-s
}

// 执行策略
func task() {
	result, err := Srv.TraderGetDetail(&model.ReqTraderDetail{TraderId: Srv.TraderId})
	if err != nil {
		fmt.Println(err)
		return
	}
	if result.Data.Status == statusStart {
		Srv.RunTimes += 1
		Run()
	} else {
		stop()
	}
}

// 停止策略
func stop() {
	c.Stop()
	os.Exit(1)
}
