package super_database

import (
	"github.com/gofrs/uuid"
	"go-hack/telegram_bot"
)

var OfferIdGenerator = uuid.DefaultGenerator
var OffersHistory []telegram_bot.StocksOffer
