package yahoo

import (
	"log"

	qt "github.com/piquette/finance-go/quote"

	"github.com/artnoi43/fngobot/adapter/fetch/common"
	"github.com/artnoi43/fngobot/usecase"
)

// Get just wraps qt.Get
func (f *fetcher) Get(tick string) (usecase.Quoter, error) {
	var q quote
	_q, err := qt.Get(tick)
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
