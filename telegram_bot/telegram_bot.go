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
		text = "❌Вы отменили действие"
	}

	actionObject := UserAction{OfferID, action}

	price := ExecuteOrder(actionObject)

	text = message.Text + "\n" + text + fmt.Sprintf("\n\nВаш баланс: %.2f\nАкций Яндекса в портфеле: %d (%.2f)\nСуммарная стоимость портфеля: %.2f",
		Balance, YNDX_amount, float64(YNDX_amount)*price, float64(YNDX_amount)*price+Balance)
	msg := tgbotapi.NewEditMessageText(message.Chat.ID, message.MessageID, text)

	_, err := Bot.botObject.Send(msg)
	if err != nil {
		println("Something went wrong")
	}
}

func ParseQueryData(queryData string) (uuid.UUID, uint8) {
	args := strings.Split(queryData, "_")

	offerID := args[0]
	action64, _ := strconv.Atoi(args[1])
	action8 := uint8(action64)

	println(offerID)
	OfferUUID, err := uuid.FromString(offerID)
	if err != nil {
		panic("Error loading Offer uuid")
	}

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
	var msg tgbotapi.MessageConfig
	var buttons []tgbotapi.InlineKeyboardButton

	//fmt.Print(offer.OfferID.String())

	switch t := offer.OfferType; t {
	case Buy:
		msg = tgbotapi.NewMessage(VanekId, BuyMessageTextFromOffer(offer))
		buttons = []tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardButtonData("✅Купить", fmt.Sprintf("%s_%d", fmt.Sprintf(offer.OfferID.String()), Accept)),
			tgbotapi.NewInlineKeyboardButtonData("❌Не покупать", fmt.Sprintf("%s_%d", fmt.Sprintf(offer.OfferID.String()), Deny))}
	case Sell:
		msg = tgbotapi.NewMessage(VanekId, SellMessageTextFromOffer(offer))
		buttons = []tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardButtonData("✅Продать", fmt.Sprintf("%s_%d", fmt.Sprintf(offer.OfferID.String()), Accept)),
			tgbotapi.NewInlineKeyboardButtonData("❌Не продавать", fmt.Sprintf("%s_%d", fmt.Sprintf(offer.OfferID.String()), Deny))}
	default:
		panic("OfferType must be Buy (0) or Sell (1)")
	}

	KeyboardMarkup := tgbotapi.NewInlineKeyboardMarkup(buttons)
	msg.ReplyMarkup = KeyboardMarkup

	if _, err := Bot.botObject.Send(msg); err != nil {
		panic(err)
	}
}

func SellMessageTextFromOffer(offer StocksOffer) string {
	if offer.OfferType != Sell {
		panic("Must be sell")
	}
	format := "Стратегия предлагает продать акцию %s в количестве %d штук по цене %.2f %s за штуку на сумму %.2f %s"
	messageText := fmt.Sprintf(format, "YNDX", offer.Amount, offer.Price, "RUB", float64(offer.Amount)*offer.Price, "RUB")
	return messageText
}

func BuyMessageTextFromOffer(offer StocksOffer) string {
	if offer.OfferType != Buy {
		panic("Must be buy")
	}
	format := "Стратегия предлагает купить акцию %s в количестве %d штук по цене %.2f %s за штуку на сумму %.2f %s"
	messageText := fmt.Sprintf(format, "YNDX", offer.Amount, offer.Price, "RUB", float64(offer.Amount)*offer.Price, "RUB")
	return messageText
}
