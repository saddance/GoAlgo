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
		fmt.Println(err, "Error while trying to run bot")
		panic(err)
	}
	bot.Debug = true

	Bot := BotWrapper{bot}

	return Bot
}

func (Bot BotWrapper) RunBot() {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := Bot.botObject.GetUpdatesChan(updateConfig)

	for update := range updates {
		Bot.HandleUpdate(update)
	}
}

func (Bot BotWrapper) HandleUpdate(update tgbotapi.Update) {
	query := update.CallbackQuery
	message := update.Message

	if query != nil {
		Bot.HandleQuery(query)
	} else if message != nil {
		Bot.EchoMessage(message)
	}

}

func (Bot BotWrapper) HandleQuery(query *tgbotapi.CallbackQuery) {
	queryData := query.Data
	args := strings.Split(queryData, "_")

	offer_id, _ := strconv.Atoi(args[0])
	action := args[1]

	fmt.Printf("User decided to %s on offer with id %d\n", action, offer_id)

	message := query.Message
	var text string
	if action == "approve" {
		text = "✅Вы подтвердили действие"
	} else if action == "deny" {
		text = "❌Вы решили не покупать акции"
	}
	msg := tgbotapi.NewEditMessageText(message.Chat.ID, message.MessageID, message.Text+"\n"+text)

	_, err := Bot.botObject.Send(msg)
	if err != nil {
		println("Something went wrong")
	}
}

func (Bot BotWrapper) EchoMessage(message *tgbotapi.Message) {
	echoMessage := tgbotapi.NewMessage(message.Chat.ID, message.Text+"\nworking as an echo bot")
	echoMessage.ReplyToMessageID = message.MessageID

	if _, err := Bot.botObject.Send(echoMessage); err != nil {
		panic(err)
	}
}

func (Bot BotWrapper) SendOffer(offer StocksOffer) {
	msg := tgbotapi.NewMessage(offer.TelegramUserId, MessageTextFromOffer(offer))

	buttons := []tgbotapi.InlineKeyboardButton{
		tgbotapi.NewInlineKeyboardButtonData("✅Купить", fmt.Sprintf("%d_approve", offer.OfferId)),
		tgbotapi.NewInlineKeyboardButtonData("❌Не покупать", fmt.Sprintf("%d_deny", offer.OfferId))}

	KeyboardMarkup := tgbotapi.NewInlineKeyboardMarkup(buttons)
	msg.ReplyMarkup = KeyboardMarkup

	if _, err := Bot.botObject.Send(msg); err != nil {
		panic(err)
	}
}

func MessageTextFromOffer(offer StocksOffer) string {
	format := "Стратегия предлагает купить акцию %s в количестве %d штук по цене %.2f %s за штуку на сумму %.2f %s"
	messageText := fmt.Sprintf(format, offer.Ticket, offer.Amount, offer.Price, offer.Currency, offer.TotalPrice, offer.Currency)
	return messageText
}
