package telegram_bot

import (
	"fmt"
	"github.com/gofrs/uuid"
	"time"
)

const (
	BOT_TOKEN = "6721745889:AAE3eqJMlp06VuoQQjDO6NpdXJEXuuzoZqY"
)

var global_actions_counter int16

func ReceiveUserAction(action UserAction) {
	println(global_actions_counter)
	global_actions_counter += 1

	fmt.Printf("Action %d) -------User decided to %s on offer with id %s--------\n", global_actions_counter, action.Action, action.OfferID)
}

func UseBot() {
	Bot := NewBot(BOT_TOKEN, false)

	Uuid, _ := uuid.NewV1()

	offer := StocksOffer{
		Uuid,
		409733921,
		"GAZP",
		4,
		100,
		"RUB",
		400,
	}

	go Bot.RunBot()
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		Bot.SendOffer(offer)
		time.Sleep(30 * time.Second)
	}
}

func main() {
	UseBot()
}
