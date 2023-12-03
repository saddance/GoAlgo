package models

type Candle struct {
	Open   float64 `xml:"open,attr"`
	Close  float64 `xml:"close,attr"`
	High   float64 `xml:"high,attr"`
	Low    float64 `xml:"low,attr"`
	Value  float64 `xml:"value,attr"`
	Volume int     `xml:"volume,attr"`
	Begin  string  `xml:"begin,attr"`
	End    string  `xml:"end,attr"`
}
