package candles

import (
	candle "go-hack/moex/candles/models"

	xmlDoc "go-hack/xml_parser/models"
)

func DocToCandles(document xmlDoc.Document) []candle.Candle {
	result := make([]candle.Candle, len(document.Data.Rows))
	for i, row := range document.Data.Rows {
		result[i] = candle.Candle{
			Open:   row.Open,
			Close:  row.Close,
			High:   row.High,
			Low:    row.Low,
			Value:  row.Value,
			Volume: row.Volume,
			Begin:  row.Begin,
			End:    row.End,
		}
	}
	return result
}
