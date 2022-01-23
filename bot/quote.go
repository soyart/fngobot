package bot

import (
	"fmt"
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

var quoteFunc = map[enums.Src]fetch.FetchFunc{
	enums.Yahoo:       yh.Get,
	enums.YahooCrypto: ct.Get,
	enums.Satang:      st.Get,
	enums.Bitkub:      bk.Get,
	enums.Binance:     bn.Get,
	enums.Coinbase:    cb.Get,
}

func (s *Security) Quote() (q fetch.Quoter, err error) {
	if s.Src.IsValid() {
		s.Tick = strings.ToUpper(s.Tick)
		q, err = quoteFunc[s.Src](s.Tick)
		if err != nil {
			return nil, err
		}
		return q, nil
	}
	return nil, fmt.Errorf(
		"invalis source %s", s.Src,
	)
}
