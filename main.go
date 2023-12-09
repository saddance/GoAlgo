package main

import (
	"fmt"
	"go-hack/moex/candles"
	"go-hack/strategy"
	"go-hack/telegram_bot"
	"time"
)

func main() {
	go telegram_bot.Bot.RunBot()
	date, err := time.Parse(time.DateOnly, "2023-12-01")
	if err != nil {
		panic(err)
	}

	res, err := candles.GetCandles("stock", "shares", "SBER", 500, date)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 5; i++ {
		strategy.HandleCandleEvent(res[i])
	}

	fmt.Println(len(res))

	fmt.Println(res)
	//telegram_bot.TestSendOffers()
}
