package telegram_bot

import (
	"fmt"
	"github.com/gofrs/uuid"
	"time"
)

const (
	BotToken = "6721745889:AAE3eqJMlp06VuoQQjDO6NpdXJEXuuzoZqY"
	VanekId  = 409733921
)

var global_actions_counter int16
var Bot = NewBot(BotToken, false)

func ReceiveUserAction(action UserAction) {
	println(global_actions_counter)
	global_actions_counter += 1

	fmt.Printf("Action %d) -------User decided to %s on offer with id %s--------\n", global_actions_counter, action.Action, action.OfferID)
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
