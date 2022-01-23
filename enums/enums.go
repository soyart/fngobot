package enums

type Src string
type QuoteType string
type Condition string

// quote sources - when adding new sources,
// also add them to validSrc below
const (
	Yahoo       Src = "yahoo"
	YahooCrypto Src = "yahooCrypto"
	Satang      Src = "satang"
	Bitkub      Src = "bitkub"
	Binance     Src = "binance"
	Coinbase    Src = "coinbase"
)

var validSrc = [6]Src{
	Yahoo,
	YahooCrypto,
	Satang,
	Bitkub,
	Binance,
	Coinbase,
}

// quote types
const (
	Bid  QuoteType = "bid"
	Ask  QuoteType = "ask"
	Last QuoteType = "last"
)

// alert condition
const (
	Lt Condition = "lt"
	Gt Condition = "gt"
)

func (s Src) IsValid() bool {
	for _, valid := range validSrc {
		if s == valid {
			return true
		}
	}
	return false
}
