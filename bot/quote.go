package bot

import (
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
	var err error
	switch s.Src {
	case enums.Yahoo:
		q, err = yh.Get(s.Tick)
	case enums.YahooCrypto:
		q, err = yh.GetCrypto(s.Tick)
	case enums.Satang:
		q, err = st.Get(s.Tick)
	case enums.Bitkub:
		q, err = bk.Get(s.Tick)
	}
	if err != nil {
		return nil, err
	}
	return q, nil
}
