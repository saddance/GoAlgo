package main

import (
	"fmt"
	"go-hack/moex/candles"
	"time"
)

func main() {
	//RunBot()

	for i := 0; i < 1; i++ {
		date, err := time.Parse(time.DateOnly, "2023-12-08")
		if err != nil {
			panic(err)
		}
		res, err := candles.GetCandles("stock", "shares", "YNDX", 500, date)
		if err != nil {
			panic(err)
		}

		for i := 0; i < len(res); i++ {
			//virtualExchange.SaveCandle(res[i])
		}
		fmt.Println(len(res))

		fmt.Println(res)
	}
}
