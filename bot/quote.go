package bot

import (
	"strings"

	"github.com/artnoi43/fngobot/enums"
	"github.com/artnoi43/fngobot/fetch"
	bk "github.com/artnoi43/fngobot/fetch/bitkub"
	ct "github.com/artnoi43/fngobot/fetch/crypto"
	st "github.com/artnoi43/fngobot/fetch/satang"
	yh "github.com/artnoi43/fngobot/fetch/yahoo"
)

func (s *Security) Quote() (fetch.Quoter, error) {
	var q fetch.Quoter
	ticker := strings.ToUpper(s.Tick)
	var err error

	switch s.Src {
	case enums.Yahoo:
		q, err = yh.Get(ticker)
	case enums.YahooCrypto:
		q, err = ct.Get(ticker)
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
