package main

import (
	"fmt"
	"github.com/bitxx/bitesla-strategy-exec/model"
)

// MACD实现
var (
	macdIndicator *Macd
	preDiff       = 0.0
	preDea        = 0.0
)

func Init() {
	fmt.Println("参数初始化")
	macdIndicator = NewMacd(12, 26, 9)

}

func Run() {
	delay := int64(5) //策略执行5次后，开始使用macd
	fmt.Println("当前次数：", Srv.RunTimes)
	if Srv.RunTimes <= delay {
		return
	}

	result, err := Srv.ExchangeGetKlineRecords(&model.ReqExchangeKline{ExchangeId: Srv.ExchangeId, Size: 1, Period: 0, CurrencyPair: "EOS_USDT"})
	if err != nil || len(result.Data) != 1 {
		fmt.Println("获取k线信息错误:", err, ",错误描述：", result.Msg)
		return
	}
	currentKline := result.Data[0]
	diff, dea, bar := macdIndicator.Update(currentKline.Close)
	//vv := currentKline.Close > currentKline.Open && ema5.Update(currentKline.Close) > ema10.Update(currentKline.Close) && ema10.Update(currentKline.Close) > ema30.Update(currentKline.Close) && diff>dea&&

	//此处逻辑需要自行调整
	if diff-dea > 0 && preDiff-preDea < 0 {
		//运行10次后，再考虑买入
		fmt.Println("全仓买入")
		fmt.Println("------------preDiff-preDea:", preDiff-preDea)
		fmt.Println("------------diff-dea:", diff-dea)
		fmt.Println("------------close:", currentKline.Close)
		fmt.Println("------------diff:", diff)
		fmt.Println("------------dea:", dea)
		fmt.Println("------------bar:", bar)
	}

	if diff-dea < 0 && preDiff-preDea > 0 {
		fmt.Println("全仓卖出")
		fmt.Println("------------preDiff-preDea:", preDiff-preDea)
		fmt.Println("------------diff-dea:", diff-dea)
		fmt.Println("------------close:", currentKline.Close)
		fmt.Println("------------diff:", diff)
		fmt.Println("------------dea:", dea)
		fmt.Println("------------bar:", bar)
	}
	preDiff = diff
	preDea = dea
}

type Macd struct {
	short  *Ema
	long   *Ema
	signal *Ema
	diff   float64
	dea    float64
	bar    float64
}

func NewMacd(short, long, signal int32) *Macd {
	return &Macd{short: NewEma(short), long: NewEma(long), signal: NewEma(signal)}
}

func (m *Macd) Update(price float64) (float64, float64, float64) {
	s := m.short.Update(price)
	l := m.long.Update(price)
	m.diff = s - l
	m.dea = m.signal.Update(m.diff)
	m.bar = 2.0 * (m.diff - m.dea)

	return m.diff, m.dea, m.bar
}

func (m *Macd) Clone() *Macd {
	return &Macd{short: m.short.Clone(), long: m.long.Clone(), signal: m.signal.Clone(), diff: m.diff, dea: m.dea, bar: m.bar}
}

type Ema struct {
	Weight float64
	result float64
	age    uint32
}

func NewEma(weight int32) *Ema {
	return &Ema{Weight: float64(weight)}
}

func (this *Ema) Update(price float64) float64 {
	if this.age == 0 {
		this.result = price
	} else {
		this.result = (2.0*price + (this.Weight-1.0)*this.result) / (this.Weight + 1.0)
	}
	this.age += 1
	return this.result
}

func (this *Ema) Clone() *Ema {
	return &Ema{Weight: this.Weight, result: this.result, age: this.age}

}
