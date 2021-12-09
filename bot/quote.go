package bot

import (
	"github.com/artnoi43/fngobot/enums"
	bk "github.com/artnoi43/fngobot/fetch/bitkub"
	st "github.com/artnoi43/fngobot/fetch/satang"
	yh "github.com/artnoi43/fngobot/fetch/yahoo"
)

// Security is a struct storing info about how to get the quotes.
type Security struct {
	Tick string
	Src  int
}

// Quoter interface specificies just 3 methods.
// Errors get returned if the sources don't support the quote type(s).
type Quoter interface {
	Last() (float64, error)
	Bid() (float64, error)
	Ask() (float64, error)
}

func (s *Security) Quote() (Quoter, error) {
	var q Quoter
	var err error
	switch s.Src {
	case enums.Yahoo:
		q, err = yh.Get(s.Tick)
		if err != nil {
			return nil, err
		}
	case enums.YahooCrypto:
		q, err = yh.GetCrypto(s.Tick)
		if err != nil {
			return nil, errYahooCryptoBidAsk
		}
	case enums.Satang:
		q, err = st.Get(s.Tick)
		if err != nil {
			return nil, errSatangLastPrice
		}
	case enums.Bitkub:
		q, err = bk.Get(s.Tick)
		if err != nil {
			return nil, err
		}
	}
	return q, nil
}
