package yahoocrypto

import (
	"log"

	"github.com/piquette/finance-go/crypto"

	"github.com/artnoi43/fngobot/adapter/fetch/common"
	"github.com/artnoi43/fngobot/usecase"
)

func (f *fetcher) Get(tick string) (usecase.Quoter, error) {
	var q quote
	_q, err := crypto.Get(tick)
	if err != nil {
		return nil, err
	}
	if _q == nil {
		log.Printf("%s not found from Yahoo Finance", tick)
		return nil, common.ErrNotFound
	}

	q.last = _q.RegularMarketPrice
	q.bid = _q.Bid
	q.ask = _q.Ask
	return &q, nil
}
