package bot

import (
	"strings"

	"github.com/artnoi43/fngobot/enums"
	bk "github.com/artnoi43/fngobot/fetch/bitkub"
	st "github.com/artnoi43/fngobot/fetch/satang"
	yh "github.com/artnoi43/fngobot/fetch/yahoo"
)

type quoter interface {
	Last() (float64, error)
	Bid() (float64, error)
	Ask() (float64, error)
}

func (s *Security) Quote() (quoter, error) {
	var q quoter
	ticker := strings.ToUpper(s.Tick)
	var err error
	switch s.Src {
	case enums.Yahoo:
		q, err = yh.Get(ticker)
	case enums.YahooCrypto:
		q, err = yh.GetCrypto(ticker)
	case enums.Satang:
		q, err = st.Get(ticker)
	case enums.Bitkub:
		q, err = bk.Get(ticker)
	}
	if err != nil {
		return nil, err
	}
	return q, nil
}
