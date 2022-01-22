package enums

type Src string
type QuoteType string
type Condition string

// quote sources
const (
	Yahoo       Src = "yahoo"
	YahooCrypto Src = "yahooCrypto"
	Satang      Src = "satang"
	Bitkub      Src = "bitkub"
	Binance     Src = "binance"
	Coinbase    Src = "coinbase"
)

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
