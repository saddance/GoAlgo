package telegram_bot

import (
	"github.com/gofrs/uuid"
)

const (
	Deny uint8 = iota
	Accept
)

type StocksOffer struct {
	OfferID        uuid.UUID
	TelegramUserId int64
	Ticket         string
	Amount         uint64
	Price          float64
	Currency       string
	TotalPrice     float64
}

type UserAction struct {
	OfferID uuid.UUID
	Action  uint8
}
