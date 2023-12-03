package models

import "encoding/xml"

type Document struct {
	XMLName xml.Name `xml:"document"`
	Data    Data     `xml:"data"`
}

type Data struct {
	ID       string   `xml:"id,attr"`
	Metadata Metadata `xml:"metadata"`
	Rows     []Row    `xml:"rows>row"`
}

type Metadata struct {
	Columns []Column `xml:"columns>column"`
}

type Column struct {
	Name    string `xml:"name,attr"`
	Type    string `xml:"type,attr"`
	Bytes   string `xml:"bytes,attr"`
	MaxSize string `xml:"max_size,attr"`
}

type Row struct {
	Open   float64 `xml:"open,attr"`
	Close  float64 `xml:"close,attr"`
	High   float64 `xml:"high,attr"`
	Low    float64 `xml:"low,attr"`
	Value  float64 `xml:"value,attr"`
	Volume int     `xml:"volume,attr"`
	Begin  string  `xml:"begin,attr"`
	End    string  `xml:"end,attr"`
}
