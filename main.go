package main

import (
	"errors"
	"github.com/jason-wj/bitesla-strategy-exec/service"
	"os"
	"strconv"
)

//测试命令： go run *.go 1 2358885120275906 eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImlkZWFfd2pAMTYzLmNvbSIsInBhc3N3b3JkIjoic2phMTIzIiwidXNlcklkIjoxLCJleHAiOjE1NTA1MDQ0MDAsImlzcyI6Ind1aiJ9.iPOoWvDbhN9HzSe2cLQt_2S8YkjD1lJJ5yKTBRrmRHA cf998e2f-ed32b3f8-42c02442-1011e 91e6e532-3bb37cad-c0a689ca-75ee5

var Srv *service.Service

func init() {
	Init()
}

func main() {
	if len(os.Args) != 6 {
		panic(errors.New("参数个数错误"))
	}
	//当前用户id 1
	currentUserIdStr := os.Args[1]
	//当前交易所 2358885120275906
	exchangeIdStr := os.Args[2]
	//令牌
	token := os.Args[3]
	//密钥key cf998e2f-ed32b3f8-42c02442-1011e
	apiKey := os.Args[4]
	//密钥secret 91e6e532-3bb37cad-c0a689ca-75ee5
	apiSecret := os.Args[5]

	currentUserId, _ := strconv.ParseInt(currentUserIdStr, 10, 64)
	exchangeId, _ := strconv.ParseInt(exchangeIdStr, 10, 64)

	//初始化服务
	Srv = &service.Service{ApiKey: apiKey, ApiSecret: apiSecret, CurrentUserId: currentUserId, Token: token, ExchangeId: exchangeId}

	Run()
}
