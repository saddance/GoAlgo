package main

import (
	"go-hack/moex/candles"
	"go-hack/strategy"
	"go-hack/telegram_bot"
	"time"
)

func main() {
	go telegram_bot.Bot.RunBot()
	//for {
	date, err := time.Parse(time.DateOnly, "2023-12-01")
	if err != nil {
		panic(err)
	}

	res, err := candles.GetCandles("stock", "shares", "YNDX", 500, date)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(res); i++ {
		go strategy.HandleCandleEvent(res[i])
		println(i)
		time.Sleep(100 * time.Millisecond)
	}
	//fmt.Println(len(res))
	//
	//fmt.Println(res)
	//	time.Sleep(30 * time.Second)
	//}
	//telegram_bot.TestSendOffers()
}
