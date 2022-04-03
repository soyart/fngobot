package enums

import "fmt"

type Src string

// quote sources - when adding new sources,
// also add them to validSrc below
const (
	Yahoo       Src = "Yahoo"
	YahooCrypto Src = "YahooCrypto"
	Satang      Src = "Satang"
	Bitkub      Src = "Bitkub"
	Binance     Src = "Binance"
	Coinbase    Src = "Coinbase"
)

var validSrc = [6]Src{
	Yahoo,
	YahooCrypto,
	Satang,
	Bitkub,
	Binance,
	Coinbase,
}

var ErrInvalidSrc = fmt.Errorf("invalid source")

func (s Src) IsValid() bool {
	for _, valid := range validSrc {
		if s == valid {
			return true
		}
	}
	return false
}
