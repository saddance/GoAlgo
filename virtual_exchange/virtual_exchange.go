package main

import (
	candle "go-hack/moex/candles/models"
	"go-hack/super_database"
	"go-hack/telegram_bot"
	"golang.org/x/exp/slices"
)

// Account represents a virtual trading account.
//type Account struct {
//	Balance  float64
//	Holdings float64 // Amount of asset currently held
//}

// create field for all candles
var candles []candle.Candle

func SaveCandle(c candle.Candle) {
	candles = append(candles, c)
}

// ExecuteOrder simulates executing a trade order.
func ExecuteOrder(action telegram_bot.UserAction) {
	offerID := action.OfferID
	offerIdInArray := slices.IndexFunc(super_database.OffersHistory, func(offer telegram_bot.StocksOffer) bool { return offer.OfferID == offerID })

	offer := super_database.OffersHistory[offerIdInArray]

	if offer.OfferType == telegram_bot.Buy {
		// Simplified: Buy asset, update balance and holdings
		super_database.Balance -= float64(offer.Amount) * offer.Price
		super_database.YNDX_amount += offer.Amount
	} else if offer.OfferType == telegram_bot.Sell {
		// Simplified: Sell asset, update balance and holdings
		super_database.Balance += float64(offer.Amount) * offer.Price
		super_database.YNDX_amount -= offer.Amount
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
