package strategy

import (
	"fmt"
	movingaverage "github.com/RobinUS2/golang-moving-average"
	"go-hack/moex/candles/models"
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

	price := getCandleValue(candle, Open)
	if prevShortAvg < prevLongAvg && curShortAvg > curLongAvg {
		//buy offer
		fmt.Print("buy\n")
		amount := int64(5)
		offerId, _ := telegram_bot.OfferIdGenerator.NewV1()

		offer := telegram_bot.StocksOffer{
			offerId,
			telegram_bot.VanekId,
			telegram_bot.Buy,
			amount,
			price,
		}

		//print(offer.Price)
		//println(offer.Amount)

		telegram_bot.SaveOffer(offer)

		telegram_bot.Bot.SendOffer(offer)
	} else if prevShortAvg > prevLongAvg && curShortAvg < curLongAvg {
		//sell all
		fmt.Print("sell\n")
		amount := int64(5)
		if telegram_bot.YNDX_amount <= amount {
			return
		}

		offerId, _ := telegram_bot.OfferIdGenerator.NewV1()

		offer := telegram_bot.StocksOffer{
			offerId,
			telegram_bot.VanekId,
			telegram_bot.Sell,
			amount,
			price,
		}

		//print(offer.Price)
		//println(offer.Amount)

		telegram_bot.SaveOffer(offer)
		telegram_bot.Bot.SendOffer(offer)
	}

	//fmt.Print(candle)
	//println(price)
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
