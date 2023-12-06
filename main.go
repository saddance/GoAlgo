// Файл для тестирования тг бота, которого я сделал
// В офере захардкожен мой айдишник, поэтому не запускайте плз этот код
// уведомления будут лететь мне

package main

import (
	"fmt"
	"go-hack/telegram_bot"
	"time"
)

const (
	BOT_TOKEN = "6721745889:AAE3eqJMlp06VuoQQjDO6NpdXJEXuuzoZqY"
)

func main() {
	Bot := telegram_bot.NewBot(BOT_TOKEN)

	offer := telegram_bot.StocksOffer{
		1,
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
		time.Sleep(3 * time.Second)
	}
}
