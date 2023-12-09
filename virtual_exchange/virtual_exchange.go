package main

import (
	candle "go-hack/moex/candles/models"
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
}

// ExecuteOrder simulates executing a trade order.
func (a *Account) ExecuteOrder(price, amount float64, buy bool) {
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

//func main() {
//	account := Account{Balance: 1000.0, Holdings: 0.0}
//	price := 50.0                             // Example price
//	amount := 5.0                             // Example amount to buy/sell
//	account.ExecuteOrder(price, amount, true) // Buy
//	fmt.Println("Account after buying:", account)
//	account.ExecuteOrder(price, amount, false) // Sell
//	fmt.Println("Account after selling:", account)
//}
