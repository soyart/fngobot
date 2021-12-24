package enums

type Src string
type QuoteType string
type Condition string

const (
	// Quote sources
	Yahoo       Src = "yahoo"
	YahooCrypto Src = "yahooCrypto"
	Satang      Src = "satang"
	Bitkub      Src = "bitkub"
	// Quote types
	Bid  QuoteType = "bid"
	Ask  QuoteType = "ask"
	Last QuoteType = "last"
	// Match conditions
	Lt Condition = "lt"
	Gt Condition = "gt"
)
