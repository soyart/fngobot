package yahoo

import (
	"log"

	qt "github.com/piquette/finance-go/quote"

	"github.com/artnoi43/fngobot/fetch"
)

// Quote for Yahoo! Finance
type quote struct {
	last float64
	bid  float64
	ask  float64
}

func (q *quote) Last() (float64, error) {
	return q.last, nil
}
func (q *quote) Bid() (float64, error) {
	return q.bid, nil
}
func (q *quote) Ask() (float64, error) {
	return q.ask, nil
}

// Get just wraps qt.Get
func Get(tick string) (fetch.Quoter, error) {
	var q quote
	_q, err := qt.Get(tick)
	if err != nil {
		return nil, err
	}
	if _q == nil {
		log.Printf("%s not found from Yahoo Finance", tick)
		return nil, fetch.ErrNotFound
	}

	q.last = _q.RegularMarketPrice
	q.bid = _q.Bid
	q.ask = _q.Ask

	return &q, nil
}
