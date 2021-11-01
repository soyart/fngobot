package bot

import (
	"github.com/artnoi43/fngobot/enums"
	bk "github.com/artnoi43/fngobot/fetch/bitkub"
	st "github.com/artnoi43/fngobot/fetch/satang"
	fn "github.com/piquette/finance-go"
	ct "github.com/piquette/finance-go/crypto"
	qt "github.com/piquette/finance-go/quote"
)

func (s *Security) Quote() (*Quote, error) {
	var q Quote
	switch s.Src {
	case enums.Yahoo:
		r, err := getQuote(s.Tick)
		if err != nil {
			return nil, err
		}
		q.Bid = r.Bid
		q.Ask = r.Ask
		q.Last = r.RegularMarketPrice
	case enums.YahooCrypto:
		r, err := getCrypto(s.Tick)
		if err != nil {
			return nil, err
		}
		q.Last = r.RegularMarketPrice
	case enums.Satang:
		r, err := getSatang(s.Tick)
		if err != nil {
			return nil, err
		}
		q.Bid = r.Bid
		q.Ask = r.Ask
	case enums.Bitkub:
		r, err := getBitkub(s.Tick)
		if err != nil {
			return nil, err
		}
		q.Bid = r.Bid
		q.Ask = r.Ask
		q.Last = r.Last
	}
	return &q, nil
}

func getQuote(tick string) (*fn.Quote, error) {
	if q, err := qt.Get(tick); err != nil {
		return nil, err

	} else if q.RegularMarketPrice == 0 {
		return nil, yahooError

	} else {
		/* Caller will usually use q.Bid or q.Ask */
		return q, nil
	}
}

func getCrypto(tick string) (*fn.CryptoPair, error) {
	if c, err := ct.Get(tick); err != nil {
		return nil, err

	} else if c.RegularMarketPrice == 0 {
		return nil, yahooError

	} else {
		/* fn.CryptoPair also contains fn.Quote */
		return c, nil
	}
}

func getSatang(tick string) (*st.Quote, error) {
	if q, err := st.Get(tick); err != nil {
		return nil, err

	} else if q.Bid == 0 {
		return nil, err

	} else {
		/* No last price from Satang API,
		 * only bid and ask prices in st.Quote. */
		return q, nil
	}
}

func getBitkub(tick string) (*bk.Quote, error) {
	if q, err := bk.Get(tick); err != nil {
		return nil, err
	} else if q.Last == 0 {
		return nil, err
	} else {
		/* Bitkub API provides many other fields,
		 * but only last, bid, and ask in *bk.Quote. */
		return q, nil
	}
}
