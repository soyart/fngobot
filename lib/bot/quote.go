package bot

import (
	"strings"

	"github.com/artnoi43/fngobot/lib/enums"
	"github.com/artnoi43/fngobot/lib/fetch"
	bn "github.com/artnoi43/fngobot/lib/fetch/binance"
	bk "github.com/artnoi43/fngobot/lib/fetch/bitkub"
	cb "github.com/artnoi43/fngobot/lib/fetch/coinbase"
	ct "github.com/artnoi43/fngobot/lib/fetch/crypto"
	st "github.com/artnoi43/fngobot/lib/fetch/satang"
	yh "github.com/artnoi43/fngobot/lib/fetch/yahoo"
)

var quoteFuncs = map[enums.Src]fetch.FetchFunc{
	enums.Yahoo:       yh.Get,
	enums.YahooCrypto: ct.Get,
	enums.Satang:      st.Get,
	enums.Bitkub:      bk.Get,
	enums.Binance:     bn.Get,
	enums.Coinbase:    cb.Get,
}

// Quote quotes Security instance. Quote sources are identified by s.Src.
// This function is called by many almost all handlers.
func (s *Security) Quote() (q fetch.Quoter, err error) {
	if s.Src.IsValid() {
		s.Tick = strings.ToUpper(s.Tick)
		quoteFunc, ok := quoteFuncs[s.Src]
		if !ok {
			// Probably forgot to add new source to quoteFีืuncs
			return nil, enums.ErrInvalidSrc
		}
		q, err = quoteFunc(s.Tick)
		if err != nil {
			return nil, err
		}
		return q, nil
	}
	// Should not happen
	// since parsing defaults to Yahoo Finance
	return nil, enums.ErrInvalidSrc
}
