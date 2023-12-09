package telegram_bot

import (
	"github.com/gofrs/uuid"
)

var OfferIdGenerator = uuid.DefaultGenerator
var OffersHistory []StocksOffer

func SaveOffer(offer StocksOffer) {
	OffersHistory = append(OffersHistory, offer)
}

func LoadOffers() []StocksOffer {
	return OffersHistory
}
