package main

import (
	"fmt"
	"github.com/jason-wj/bitesla-strategy-exec/model"
)

//MACD实现

func Init() {
	fmt.Println("参数初始化")
}

func Run() {
	result, err := Srv.ExchangeGetKlineRecords(&model.ReqExchangeKline{ExchangeId: Srv.ExchangeId, Size: 1, Period: 4, CurrencyPair: "ETH_USDT"})
	if err != nil || len(result.Data) != 1 {
		fmt.Println("获取k线信息错误:", err)
		return
	}
	macdIndicator := NewMacd(12, 26, 9)
	f := result.Data[0].Close
	diff, dea, macd := macdIndicator.Update(f)
	fmt.Println("------------close:", f)
	fmt.Println("------------diff:", diff)
	fmt.Println("------------dea:", dea)
	fmt.Println("------------macd:", macd)

}

type Macd struct {
	short  *Ema
	long   *Ema
	signal *Ema
	diff   float64
	dea    float64
	macd   float64
}

func NewMacd(short, long, signal int32) *Macd {
	return &Macd{short: NewEma(short), long: NewEma(long), signal: NewEma(signal)}
}

func (m *Macd) Update(price float64) (float64, float64, float64) {
	s := m.short.Update(price)
	l := m.long.Update(price)
	m.diff = s - l
	m.dea = m.signal.Update(m.diff)
	m.macd = 2.0 * (m.diff - m.dea)

	return m.diff, m.dea, m.macd
}

func (m *Macd) Clone() *Macd {
	return &Macd{short: m.short.Clone(), long: m.long.Clone(), signal: m.signal.Clone(), diff: m.diff, dea: m.dea, macd: m.macd}
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
