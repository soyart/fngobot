package yahoo

import (
	"log"

	qt "github.com/piquette/finance-go/quote"

	"github.com/artnoi43/fngobot/adapter/fetch/common"
	"github.com/artnoi43/fngobot/usecase"
)

// Get just wraps qt.Get
func (f *fetcher) Get(tick string) (usecase.Quoter, error) {
	var q common.Quote
	_q, err := qt.Get(tick)
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

	return &q, nil
}
