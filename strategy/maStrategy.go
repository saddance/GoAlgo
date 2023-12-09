package strategy

import (
	movingaverage "github.com/RobinUS2/golang-moving-average"
	"go-hack/moex/candles/models"
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

func handleCandleEvent(candle models.Candle) {
	prevShortAvg := shortMA.Avg()
	prevLongAvg := longMA.Avg()

	shortMA.Add(getCandleValue(candle, Close))
	longMA.Add(getCandleValue(candle, Low))
	curShortAvg := shortMA.Avg()
	curLongAvg := longMA.Avg()

	if prevShortAvg < prevLongAvg && curShortAvg > curLongAvg {
		//buy
	} else if prevShortAvg > prevLongAvg && curShortAvg < curLongAvg {
		//sell all
	}
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
