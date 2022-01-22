package bot

import (
	"strings"

	"github.com/artnoi43/fngobot/enums"
	"github.com/artnoi43/fngobot/fetch"
	bn "github.com/artnoi43/fngobot/fetch/binance"
	bk "github.com/artnoi43/fngobot/fetch/bitkub"
	cb "github.com/artnoi43/fngobot/fetch/coinbase"
	ct "github.com/artnoi43/fngobot/fetch/crypto"
	st "github.com/artnoi43/fngobot/fetch/satang"
	yh "github.com/artnoi43/fngobot/fetch/yahoo"
)

func (s *Security) Quote() (q fetch.Quoter, err error) {
	s.Tick = strings.ToUpper(s.Tick)
	switch s.Src {
	case enums.Yahoo:
		q, err = yh.Get(s.Tick)
	case enums.YahooCrypto:
		q, err = ct.Get(s.Tick)
	case enums.Satang:
		q, err = st.Get(s.Tick)
	case enums.Bitkub:
		q, err = bk.Get(s.Tick)
	case enums.Binance:
		q, err = bn.Get(s.Tick)
	case enums.Coinbase:
		q, err = cb.Get(s.Tick)
	}
	return
}
