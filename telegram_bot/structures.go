package telegram_bot

//type StocksOffer struct {
//	offerId        uint64
//	telegramUserId int64
//	ticket         string
//	amount         uint64
//	price          float64
//	totalPrice     float64
//}

type StocksOffer struct {
	OfferId        uint64
	TelegramUserId int64
	Ticket         string
	Amount         uint64
	Price          float64
	Currency       string
	TotalPrice     float64
}

type UserAction struct {
	OfferId        int64
	telegramUserId int64
	action         bool
}
