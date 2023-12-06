package main

import (
	"fmt"
	"go-hack/moex/candles"
	"time"
)

func main() {
	date, err := time.Parse(time.DateOnly, "2023-12-01")
	if err != nil {
		panic(err)
	}

	res, err := candles.GetCandles("stock", "shares", "SBER", 500, date)
	if err != nil {
		panic(err)
	}

	fmt.Println(len(res))

	fmt.Println(res)
}
