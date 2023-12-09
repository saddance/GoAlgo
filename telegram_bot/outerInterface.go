package telegram_bot

import (
	"fmt"
	"github.com/gofrs/uuid"
	"go-hack/super_database"
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
	offerID := action.OfferID
	offerIdInArray := slices.IndexFunc(super_database.OffersHistory, func(offer StocksOffer) bool { return offer.OfferID == offerID })

	offer := super_database.OffersHistory[offerIdInArray]

	if offer.OfferType == Buy {
		// Simplified: Buy asset, update balance and holdings
		super_database.Balance -= float64(offer.Amount) * offer.Price
		super_database.YNDX_amount += offer.Amount
	} else if offer.OfferType == Sell {
		// Simplified: Sell asset, update balance and holdings
		super_database.Balance += float64(offer.Amount) * offer.Price
		super_database.YNDX_amount -= offer.Amount
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
