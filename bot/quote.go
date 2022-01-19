package bot

import (
	"github.com/artnoi43/fngobot/enums"
	"github.com/artnoi43/fngobot/fetch"
	bk "github.com/artnoi43/fngobot/fetch/bitkub"
	ct "github.com/artnoi43/fngobot/fetch/crypto"
	st "github.com/artnoi43/fngobot/fetch/satang"
	yh "github.com/artnoi43/fngobot/fetch/yahoo"
)

func (s *Security) Quote() (fetch.Quoter, error) {
	var q fetch.Quoter
	var err error

	switch s.Src {
	case enums.Yahoo:
		q, err = yh.Get(s.Tick)
	case enums.YahooCrypto:
		q, err = ct.Get(s.Tick)
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
