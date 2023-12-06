// Файл для тестирования тг бота, которого я сделал
// В офере захардкожен мой айдишник, поэтому не запускайте плз этот код
// уведомления будут лететь мне

package main

import (
	"fmt"
	"github.com/gofrs/uuid"
	"go-hack/telegram_bot"
	"time"
)

const (
	BOT_TOKEN = "6721745889:AAE3eqJMlp06VuoQQjDO6NpdXJEXuuzoZqY"
)

func main() {
	Bot := telegram_bot.NewBot(BOT_TOKEN)

	Uuid, _ := uuid.NewV1()

	offer := telegram_bot.StocksOffer{
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
