package yahoo

import (
	"errors"
	"log"

	"github.com/artnoi43/fngobot/fetch"
	"github.com/piquette/finance-go/crypto"
)

type cryptoQuote struct {
	last float64
	bid  float64
	ask  float64
}

func (q *cryptoQuote) Last() (float64, error) {
	return q.last, nil
}
func (q *cryptoQuote) Bid() (float64, error) {
	return 0, errors.New("yahoo_crypto: bid not supported")
}
func (q *cryptoQuote) Ask() (float64, error) {
	return 0, errors.New("yahoo_crypto: ask not supported")
}

func GetCrypto(tick string) (*cryptoQuote, error) {
	var q cryptoQuote
	_q, err := crypto.Get(tick)
	if err != nil {
		return nil, err
	}
	if _q == nil {
		log.Printf("%s not found in Satang JSON", tick)
		return nil, fetch.ErrNotFound
	}

	q.last = _q.RegularMarketPrice
	q.bid = _q.Bid
	q.ask = _q.Ask
	return &q, nil
}
