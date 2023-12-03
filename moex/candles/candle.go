package candles

import (
	"fmt"
	"strconv"
	"time"

	"github.com/imroc/req/v3"

	"go-hack/moex/candles/models"
	"go-hack/xml_parser"
)

var client = req.C()

func GetCandles(engine, market, security string, start int, from time.Time) ([]models.Candle, error) {
	params := map[string]interface{}{
		"from":  from.Format(time.DateOnly),
		"till":  "2037-12-31",
		"start": strconv.Itoa(start),
	}

	url := fmt.Sprintf(
		"https://iss.moex.com/iss/engines/%s/markets/%s/securities/%s/candles",
		engine, market, security,
	)
	for key, value := range params {
		url += fmt.Sprintf("?%s=%s", key, value)
	}
	fmt.Println(url)

	resp, err := client.R().Get(url)

	if err != nil {
		return nil, err
	}

	response, err := resp.ToBytes()
	if err != nil {
		return nil, err
	}

	doc, err := xml_parser.ParseXML(response)

	return DocToCandles(doc), err
}
