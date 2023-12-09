package telegram_bot

import (
	"github.com/gofrs/uuid"
	"sync"
)

var OfferIdGenerator = uuid.DefaultGenerator
var OffersHistory []StocksOffer
var mutex sync.Mutex // Mutex to protect OffersHistory

func SaveOffer(offer StocksOffer) {
	mutex.Lock() // Lock before modifying OffersHistory
	OffersHistory = append(OffersHistory, offer)
	mutex.Unlock() // Unlock after modifying
}

func LoadOffers() []StocksOffer {
	mutex.Lock()         // Lock before accessing OffersHistory
	defer mutex.Unlock() // Unlock after accessing (defer ensures that the unlock will happen)
	return OffersHistory
}
