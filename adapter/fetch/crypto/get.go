package yahoocrypto

import (
	"log"

	"github.com/piquette/finance-go/crypto"

	"github.com/artnoi43/fngobot/adapter/fetch/common"
	"github.com/artnoi43/fngobot/internal/enums"
	"github.com/artnoi43/fngobot/usecase"
)

func (f *fetcher) Get(tick string) (usecase.Quoter, error) {
	var q common.ApiQuote
	_q, err := crypto.Get(tick)
	if err != nil {
		return nil, err
	}
	if _q == nil {
		log.Printf("%s not found from Yahoo Finance\n", tick)
		return nil, common.ErrNotFound
	}

	q.Last = _q.RegularMarketPrice
	q.Bid = _q.Bid
	q.Ask = _q.Ask
	q.Src = enums.YahooCrypto
	return &q, nil
}
