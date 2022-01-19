package bot

import (
	"github.com/artnoi43/fngobot/enums"
	"github.com/artnoi43/fngobot/fetch"
	"github.com/pkg/errors"
)

// Alert struct stores info about price alerts
type Alert struct {
	Security  `json:"security"`
	Condition enums.Condition `json:"condition"`
	QuoteType enums.QuoteType `json:"quote"`
	Target    float64         `json:"target"`
}

func GetQuoteAndAlert(
	a *Alert,
	matched chan<- bool,
	errChan chan<- error,
) {
	q, err := a.Security.Quote()
	if err != nil {
		// false is sent
		// if Quote() returns an error
		errChan <- errors.Wrapf(
			err,
			"failed to get quote for %v from %v",
			a.Security.Tick, a.Src,
		)
	}
	// @TODO: use errChan - now it's here for testing
	Match(a, matched, errChan, q)
}

// Match sends a truthy value into matched channel
// if the specified market condition is matched.
func Match(
	a *Alert,
	matched chan<- bool,
	errChan chan<- error,
	marketQuote fetch.Quoter,
) {
	var p float64
	var err error
	switch a.QuoteType {
	case enums.Bid:
		p, err = marketQuote.Bid()
	case enums.Ask:
		p, err = marketQuote.Ask()
	case enums.Last:
		p, err = marketQuote.Last()
	}

	if err != nil {
		errChan <- errors.Wrap(err, "invalid quote type")
		return
	}

	switch a.Condition {
	case enums.Lt:
		if p <= a.Target {
			matched <- true
		} else {
			matched <- false
		}
	case enums.Gt:
		if p >= a.Target {
			matched <- true
		} else {
			matched <- false
		}
	}
}
