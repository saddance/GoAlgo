package strategy

import (
	movingaverage "github.com/RobinUS2/golang-moving-average"
	"go-hack/moex/candles/models"
	"go-hack/super_database"
	"go-hack/telegram_bot"
)

type MovingType int

const (
	Open MovingType = iota
	High
	Low
	Close
)

var shortMA = movingaverage.New(5)
var longMA = movingaverage.New(10)

func HandleCandleEvent(candle models.Candle) {
	prevShortAvg := shortMA.Avg()
	prevLongAvg := longMA.Avg()

	shortMA.Add(getCandleValue(candle, Close))
	longMA.Add(getCandleValue(candle, Low))
	curShortAvg := shortMA.Avg()
	curLongAvg := longMA.Avg()

	var offer telegram_bot.StocksOffer

	price := getCandleValue(candle, Close)
	if prevShortAvg < prevLongAvg && curShortAvg > curLongAvg {
		//buy offer
		amount := uint64(5)

		offerId, _ := super_database.OfferIdGenerator.NewV1()
		offer = telegram_bot.StocksOffer{
			offerId,
			telegram_bot.VanekId,
			telegram_bot.Buy,
			amount,
			price,
		}
	} else if prevShortAvg > prevLongAvg && curShortAvg < curLongAvg {
		//sell all
		amount := super_database.YNDX_amount

		if amount != 0 {
			offerId, _ := super_database.OfferIdGenerator.NewV1()
			offer = telegram_bot.StocksOffer{
				offerId,
				telegram_bot.VanekId,
				telegram_bot.Sell,
				amount,
				price,
			}
		} else {
			return
		}
	}
	super_database.OffersHistory = append(super_database.OffersHistory, offer)
	telegram_bot.Bot.SendOffer(offer)
}

func getCandleValue(candle models.Candle, movingType MovingType) float64 {
	switch movingType {
	case Open:
		return candle.Open
	case High:
		return candle.High
	case Low:
		return candle.Low
	case Close:
		return candle.Close
	default:
		return candle.Close
	}
}
