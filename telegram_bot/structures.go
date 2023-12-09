package telegram_bot

import (
	"github.com/gofrs/uuid"
)

const (
	Deny uint8 = iota
	Accept
)

const (
	Buy uint8 = iota
	Sell
)

type StocksOffer struct {
	OfferID        uuid.UUID
	TelegramUserId int64
	OfferType      uint8
	Amount         int64
	Price          float64
}

type UserAction struct {
	OfferID uuid.UUID
	Action  uint8
}
