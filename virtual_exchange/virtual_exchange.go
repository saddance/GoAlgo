package main

import (
	candle "go-hack/moex/candles/models"
	"go-hack/strategy"
)

// Account represents a virtual trading account.
type Account struct {
	Balance  float64
	Holdings float64 // Amount of asset currently held
}

// create field for all candles
var candles []candle.Candle

func SaveCandle(c candle.Candle) {
	candles = append(candles, c)
	strategy.HandleCandleEvent(c)
}

// ExecuteOrder simulates executing a trade order.
func (a *Account) ExecuteOrder(price, amount float64, buy bool) {

	//get order from memory

	if buy {
		// Simplified: Buy asset, update balance and holdings
		a.Holdings += amount
		a.Balance -= price * amount
	} else {
		// Simplified: Sell asset, update balance and holdings
		a.Holdings -= amount
		a.Balance += price * amount
	}
}
