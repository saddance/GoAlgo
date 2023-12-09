package telegram_bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gofrs/uuid"
	"strconv"
	"strings"
)

type BotWrapper struct {
	botObject *tgbotapi.BotAPI
}

func NewBot(BotToken string, debug bool) BotWrapper {
	bot, err := tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		fmt.Println(err, "Error while trying to run bot")
		panic(err)
	}
	bot.Debug = debug

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
	OfferID, action := ParseQueryData(query.Data)

	message := query.Message
	var text string
	if action == Accept {
		text = "✅Вы подтвердили действие"
	} else if action == Deny {
		text = "❌Вы решили не покупать акции"
	}

	msg := tgbotapi.NewEditMessageText(message.Chat.ID, message.MessageID, message.Text+"\n"+text)

	_, err := Bot.botObject.Send(msg)
	if err != nil {
		println("Something went wrong")
	}

	actionObject := UserAction{OfferID, action}

	ReceiveUserAction(actionObject)
}

func ParseQueryData(queryData string) (uuid.UUID, uint8) {
	args := strings.Split(queryData, "_")

	offerID := args[0]
	action64, _ := strconv.Atoi(args[1])
	action8 := uint8(action64)

	OfferUUID, _ := uuid.FromString(offerID)

	return OfferUUID, action8
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
		tgbotapi.NewInlineKeyboardButtonData("✅Купить", fmt.Sprintf("%d_%d", offer.OfferID, Accept)),
		tgbotapi.NewInlineKeyboardButtonData("❌Не покупать", fmt.Sprintf("%d_%d", offer.OfferID, Deny))}

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
