package xml_parser

import (
	"encoding/xml"
	"go-hack/xml_parser/models"
)

func ParseXML(xmlData []byte) (models.Document, error) {
	var doc models.Document
	err := xml.Unmarshal(xmlData, &doc)
	if err != nil {
		return models.Document{}, err
	}
	return doc, nil
}
