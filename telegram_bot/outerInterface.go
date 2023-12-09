package telegram_bot

import (
	"fmt"
	"github.com/gofrs/uuid"
	"golang.org/x/exp/slices"
	"time"
)

const (
	BotToken = "6721745889:AAE3eqJMlp06VuoQQjDO6NpdXJEXuuzoZqY"
	VanekId  = 409733921
)

var global_actions_counter int16
var Bot = NewBot(BotToken, false)

func ExecuteOrder(action UserAction) {
	LocalOffersHistory := LoadOffers()
	offerID := action.OfferID
	offerIdInArray := slices.IndexFunc(LocalOffersHistory, func(offer StocksOffer) bool { return offer.OfferID == offerID })

	//fmt.Println(OffersHistory)
	//fmt.Println(offerID)
	//
	//fmt.Println("------------------------------------------- %d", offerIdInArray)

	offer := LocalOffersHistory[offerIdInArray]

	if offer.OfferType == Buy {
		// Simplified: Buy asset, update balance and holdings
		Balance -= float64(offer.Amount) * offer.Price
		YNDX_amount += offer.Amount
	} else if offer.OfferType == Sell {
		// Simplified: Sell asset, update balance and holdings
		Balance += float64(offer.Amount) * offer.Price
		YNDX_amount -= offer.Amount
	}
}

func TestSendOffers() {
	Uuid, _ := uuid.NewV1()
	offer := StocksOffer{
		Uuid,
		VanekId,
		Buy,
		4,
		100,
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
		Bot.SendOffer(offer)
		time.Sleep(5 * time.Second)
	}
}

//func main() {
//	go Bot.RunBot()
//	TestSendOffers()
//}
