package telegram_bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"strings"
)

type BotWrapper struct {
	botObject *tgbotapi.BotAPI
}

func NewBot(BotToken string) BotWrapper {
	bot, err := tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		fmt.Println(err, "Cant run bot pum-pum-pum")
		panic(err)
	}
	bot.Debug = false

	Bot := BotWrapper{bot}

	return Bot
}

func (Bot BotWrapper) RunBot() {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := Bot.botObject.GetUpdatesChan(updateConfig)

	for update := range updates {

		query := update.CallbackQuery
		if query != nil {
			queryData := query.Data
			args := strings.Split(queryData, "_")

			offer_id, _ := strconv.Atoi(args[0])
			action := args[1]

			fmt.Printf("User decided to %s on offer with id %d\n", action, offer_id)

			message := query.Message
			var text string
			if action == "approve" {
				text = "✅Вы подтвердили действие"
			} else if action == "cancel" {
				text = "❌Вы решили не покупать акции"
			}
			msg := tgbotapi.NewEditMessageText(message.Chat.ID, message.MessageID, message.Text+"\n"+text)

			_, err := Bot.botObject.Send(msg)
			if err != nil {
				println("Something went wrong")
			}
		} else if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text+"\nworking as an echo bot")
			msg.ReplyToMessageID = update.Message.MessageID

			if _, err := Bot.botObject.Send(msg); err != nil {
				panic(err)
			}
		}

	}
}

func (Bot BotWrapper) SendOffer(offer StocksOffer) {
	msg := tgbotapi.NewMessage(offer.TelegramUserId,
		fmt.Sprintf(
			"Стратегия предлагает купить акцию %s в количестве "+
				"%d штук по цене %.2f %s за штуку на сумму %.2f %s", offer.Ticket, offer.Amount, offer.Price, offer.Currency, offer.TotalPrice, offer.Currency))

	buttons := []tgbotapi.InlineKeyboardButton{tgbotapi.NewInlineKeyboardButtonData("✅Купить", fmt.Sprintf("%d_approve", offer.OfferId)),
		tgbotapi.NewInlineKeyboardButtonData("❌Не покупать", fmt.Sprintf("%d_cancel", offer.OfferId))}

	KeyboardMarkup := tgbotapi.NewInlineKeyboardMarkup(buttons)
	msg.ReplyMarkup = KeyboardMarkup

	if _, err := Bot.botObject.Send(msg); err != nil {
		panic(err)
	}
}
